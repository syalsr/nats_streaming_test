package cache

import "Wildberries_L0/detail"

type Cache struct {
	cache map[string]detail.OrderInfo
}

func (c *Cache) SaveToCache(info *detail.OrderInfo) {
	c.cache[info.OrderUID] = *info
}

func NewCache() Cache {
	return Cache{cache: make(map[string]detail.OrderInfo)}
}

func ReloadCacheFromDatabase(info *detail.OrderInfo) Cache {
	Cache := Cache{cache: make(map[string]detail.OrderInfo)}
	Cache.cache[info.OrderUID] = *info
	return Cache
}
