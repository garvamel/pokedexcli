package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	dur := interval * time.Nanosecond
	var c = Cache{
		cache: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	go c.reapLoop(dur)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cE, ok := c.cache[key]
	if ok {
		return cE.val, ok
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {

	tick := time.NewTicker(interval)

	for true {
		select {
		case <-tick.C:
			t := time.Now()
			c.mu.Lock()
			for key := range c.cache {
				if t.Sub(c.cache[key].createdAt) > interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
