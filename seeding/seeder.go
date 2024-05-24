package seeding

import (
	"fmt"

	"github.com/anubhav100rao/cache_server/cache"
	config "github.com/anubhav100rao/cache_server/config"
)

type Payload struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func SeedWithRandomData(cache *cache.Cache) {
	// Register the seeding package
	for i := 0; i < config.CACHE_DEFAULT_SIZE/2; i++ {
		key := fmt.Sprintf("key%d", i)
		data := Payload{
			Key: fmt.Sprintf("key%d", i),
			Value: map[string]interface{}{
				"name": fmt.Sprintf("name%d", i),
				"age":  i,
			},
		}

		cache.Set(key, data)
	}
}
