package pokecache

import (
	"sync" //for mutex
	"time"
)

//cacheEntry struct {
//	createdAt - time.Time that represents when
// 	the entry was created
//	val - []byte that represents raw data we're caching
//}

type cacheEntry struct {
	createdAt time.Time //time.Time is a type (struct)
	val       []byte
}

//create a Cache struct holding map[string]cacheEntry
// Cache struct {
//	map[string]cacheEntry
//	mutex to protect map
//}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex //locking for goroutines
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval) //make reapLoop method
	return &c
}

//create cache.Add() that adds new entry to the cache
//parameter: key (string) and a val ([byte])

//func (receiverName ReceiverType) MethodName(params) returnType

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

//create cache.Get() gets entry from cache
//parameters: key(string)
//return: []byte and bool - true if entry is found

func (c *Cache) Get(key string) ([]byte, bool) {
	// if c.entries[key].val == nil {
	// 	return nil, false
	// }
	// return c.entries[key].val, true
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

//create cache.reapLoop() - called when cache is created
//takes in interval
//remove entries that are older than interval
//time.Ticker can be useful

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C { //ticker.C , type <-chan Time
		c.mu.Lock()
		for key, cacheEntry := range c.entries {
			if time.Since(cacheEntry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
