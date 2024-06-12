package pokecache

import (
    "time"
    "sync"
)

type cacheEntry struct {
    createdAt time.Time
    val	      []byte
}

type Cache struct {
    entries map[string]cacheEntry
    mutex sync.Mutex
    interval time.Duration
}


func NewCache(interval time.Duration) *Cache {
    cache := &Cache{
	entries: make(map[string]cacheEntry),
	interval: interval,
    }
    go cache.reapLoop()
    return cache
}

func (c *Cache) Add(key string, val []byte) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    entry := cacheEntry{
	createdAt: time.Now(),
	val:	   val,
    }
    c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    entry, ok := c.entries[key]
    if !ok {
	return nil, false
    }
    return entry.val, true
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()

    for {
	<-ticker.C
	c.mutex.Lock()
	currentTime := time.Now()
	for key, entry := range c.entries {
	    if currentTime.Sub(entry.createdAt) > c.interval {
		delete(c.entries, key)
	    }
	}
	c.mutex.Unlock()
    }
}





