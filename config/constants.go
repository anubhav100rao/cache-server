package config

import (
	"time"

	"github.com/anubhav100rao/cache_server/eviction"
)

// CACHE_DEFAULT_SIZE is the default size of the cache
const CACHE_DEFAULT_SIZE = 10

// DURATION is the default duration for which an item is stored in the cache
const DURATION = 1000 * time.Second

type CacheConfig struct {
	Size int
	eviction.EvictionPolicy
}

func NewCacheConfig(size int, policy eviction.EvictionPolicy) *CacheConfig {
	return &CacheConfig{
		Size:           size,
		EvictionPolicy: policy,
	}
}
