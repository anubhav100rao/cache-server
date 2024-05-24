package seeding

import (
	"fmt"
	"time"

	cachepkg "github.com/anubhav100rao/cache_server/cache"
	config "github.com/anubhav100rao/cache_server/config"
)

type Payload struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func SeedWithRandomData(cache *cachepkg.Cache) {
	// Register the seeding package
	for i := 0; i < config.CACHE_DEFAULT_SIZE/2; i++ {
		key := fmt.Sprintf("key%d", i)
		expiryTime := time.Now().Add(config.DURATION)
		data := cachepkg.CacheItem{
			ExpiryTime: expiryTime,
			Value: Payload{
				Key: key,
				Value: map[string]interface{}{
					"key":   key,
					"value": fmt.Sprintf("value%d", i),
				},
			},
		}
		cache.Set(key, data, config.DURATION)
	}
}
