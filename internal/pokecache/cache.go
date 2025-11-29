package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data    map[string]CacheEntry
	mux     *sync.Mutex
	timeout time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	mux := &sync.Mutex{}
	newCache := Cache{
		timeout: interval,
		mux:     mux,
		data:    make(map[string]CacheEntry),
	}
	go newCache.ReapLoop()
	return newCache
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// ReapLoop -
func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.timeout)
	defer ticker.Stop()

	for range ticker.C {
		c.Reap(time.Now().UTC())
	}
}

// Reap -
func (c *Cache) Reap(now time.Time) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-c.timeout)) {
			delete(c.data, k)
		}
	}
}
