package cache

import (
	"Wildberries_L0/database"
	"Wildberries_L0/detail"
	"context"
	"github.com/jackc/pgx/v4"
)

type Cache struct {
	cache map[string]detail.OrderInfo
}

func (c *Cache) SaveToCache(info *detail.OrderInfo) {
	c.cache[info.OrderUID] = *info
}

func newCache() *Cache {
	return &Cache{cache: make(map[string]detail.OrderInfo)}
}

func (c *Cache) GetOrderByUID(uid string) detail.OrderInfo {
	return c.cache[uid]
}

func getOrderUID(connection *pgx.Conn) (ids []string) {
	query := `
			SELECT array_agg(order_uid) FROM orders
`
	connection.QueryRow(context.Background(), query).Scan(&ids)
	return
}

func LoadCacheFromDatabase(connection *pgx.Conn) *Cache {
	orderUIDs := getOrderUID(connection)
	orderCache := newCache()
	for _, id := range orderUIDs {
		order := database.GetOrderByUID(connection, id)
		orderCache.cache[id] = *order
	}
	return orderCache
}
