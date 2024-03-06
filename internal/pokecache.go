package internal

import (
	"sync"
	"time"
)

var mu = &sync.RWMutex{}

type Cache struct {
	cacheEntry map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	mu.Lock()
	c.cacheEntry[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	mu.RLock()
	cacheEntry, ok := c.cacheEntry[key]
	mu.RUnlock()

	if ok {
		return cacheEntry.val, true
	}

	return []byte{}, false
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]CacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(interval)
		done <- true
	}()

	for {
		select {
		case <-done:
			mu.Lock()
			c.cacheEntry = nil
			mu.Unlock()
			return
		}
	}
}
