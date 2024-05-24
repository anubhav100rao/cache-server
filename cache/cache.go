package cache

import (
	"sync"
	"time"

	eviction "github.com/anubhav100rao/cache_server/eviction"
)

type CacheItem struct {
	Value      interface{}
	ExpiryTime time.Time
}

// Cache represents the in-memory cache
type Cache struct {
	capacity       int
	evictionPolicy eviction.EvictionPolicy
	cache          map[string]CacheItem
	mutex          sync.RWMutex
}

// NewCache creates a new cache with the given capacity and eviction policy
func NewCache(capacity int, evictionPolicy eviction.EvictionPolicy) *Cache {
	return &Cache{
		capacity:       capacity,
		evictionPolicy: evictionPolicy,
		cache:          make(map[string]CacheItem),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, exists := c.cache[key]

	if !exists {
		return nil, false
	}

	if time.Now().After(value.ExpiryTime) {
		c.mutex.RUnlock()
		c.Delete(key)
		c.mutex.RLock()
		return nil, false
	}

	c.evictionPolicy.RecordAccess(key)

	return value, exists
}

// Set adds a key-value pair to the cache
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expiryTime := time.Now().Add(duration)

	if _, exists := c.cache[key]; exists {
		c.cache[key] = CacheItem{Value: value, ExpiryTime: expiryTime}
		c.evictionPolicy.RecordAccess(key)
		return
	}

	if len(c.cache) >= c.capacity {
		evictKey := c.evictionPolicy.Evict()
		if evictKey != "" {
			delete(c.cache, evictKey)
		}
	}

	c.cache[key] = CacheItem{Value: value, ExpiryTime: expiryTime}
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

	c.cache = make(map[string]CacheItem)
	c.evictionPolicy.Clear()
}

// GetAll returns all key-value pairs from the cache
func (c *Cache) GetAll() map[string]CacheItem {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.cache
}

func (c *Cache) StartCleanup(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			c.cleanup()
		}
	}()
}

func (c *Cache) cleanup() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, item := range c.cache {
		if time.Now().After(item.ExpiryTime) {
			delete(c.cache, key)
			c.evictionPolicy.Remove(key)
		}
	}
}
