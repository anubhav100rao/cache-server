package config

import "github.com/anubhav100rao/cache_server/eviction"

const CACHE_DEFAULT_SIZE = 10

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
