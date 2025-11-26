package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(url string, val []byte) {
	c.cache[url] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	result, exists :=  c.cache[url]
	if !exists {
		return nil, exists
	}
	return result.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) PrintCache() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, entry := range c.cache {
		fmt.Println("Name:", name)
		fmt.Println("Cache Entry:", entry)
	}
}