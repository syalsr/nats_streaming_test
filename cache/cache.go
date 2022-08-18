package cache

import (
	"Wildberries_L0/detail"
	"context"
	"encoding/json"
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

func (c *Cache) GetOrderByUID(uid string) (value detail.OrderInfo, inCache bool) {
	value, inCache = c.cache[uid]
	if inCache {
		return value, true
	}
	return detail.OrderInfo{}, false
}

func LoadCacheFromDatabase(connection *pgx.Conn) (orderCache *Cache) {
	orderCache = newCache()
	rows, err := connection.Query(context.Background(), "SELECT order_uid, info FROM orders")
	if err != nil {
		panic(err)
	}
	//defer rows.Close()
	for rows.Next() {
		var orderUID string
		var info string
		err := rows.Scan(&orderUID, &info)
		if err != nil {
			panic(err)
		}
		var orderInfo detail.OrderInfo
		err = json.Unmarshal([]byte(info), &orderInfo)
		if err != nil {
			panic(err)
		}
		orderCache.SaveToCache(&orderInfo)
	}
	return orderCache
}
