package cache

import (
	"sync"

	eviction "github.com/anubhav100rao/cache_server/eviction"
)

// Cache represents the in-memory cache
type Cache struct {
	capacity       int
	evictionPolicy eviction.EvictionPolicy
	cache          map[string]interface{}
	mutex          sync.RWMutex
}

// NewCache creates a new cache with the given capacity and eviction policy
func NewCache(capacity int, evictionPolicy eviction.EvictionPolicy) *Cache {
	return &Cache{
		capacity:       capacity,
		evictionPolicy: evictionPolicy,
		cache:          make(map[string]interface{}),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, exists := c.cache[key]
	if exists {
		c.evictionPolicy.RecordAccess(key)
	}
	return value, exists
}

// Set adds a key-value pair to the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.cache[key]; exists {
		c.cache[key] = value
		c.evictionPolicy.RecordAccess(key)
		return
	}

	if len(c.cache) >= c.capacity {
		evictKey := c.evictionPolicy.Evict()
		if evictKey != "" {
			delete(c.cache, evictKey)
		}
	}

	c.cache[key] = value
	c.evictionPolicy.Add(key)
}

// Delete removes a key-value pair from the cache
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.cache[key]; exists {
		delete(c.cache, key)
		c.evictionPolicy.Remove(key)
	}
}

// Clear removes all key-value pairs from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache = make(map[string]interface{})
	c.evictionPolicy.Clear()
}

// GetAll returns all key-value pairs from the cache
func (c *Cache) GetAll() map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.cache
}
