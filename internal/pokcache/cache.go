package pokcache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	mu       sync.Mutex
	cacheMap map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.cacheMap[key]
	if !found {
		return nil, false
	}

	return entry.val, true

}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.cacheMap {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.cacheMap, k)
			}
		}
		c.mu.Unlock()
	}
}
