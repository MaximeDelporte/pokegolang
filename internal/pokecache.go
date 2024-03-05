package internal

import "time"

type Cache struct {
	cacheEntry map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, value []byte) {

}

func (c Cache) Get(key string) ([]byte, bool) {
	return []byte{}, false
}

func NewCache(interval time.Duration) Cache {
	return Cache{}
}
