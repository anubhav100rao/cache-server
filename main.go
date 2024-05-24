package main

import (
	"fmt"

	"net/http"

	cache "github.com/anubhav100rao/cache_server/cache"
	"github.com/anubhav100rao/cache_server/config"
	eviction "github.com/anubhav100rao/cache_server/eviction"
	seeding "github.com/anubhav100rao/cache_server/seeding"
	service "github.com/anubhav100rao/cache_server/service"
)

func main() {
	// Use a custom eviction policy, here we use RandomEvictionPolicy as an example
	randomPolicy := eviction.NewRandomEvictionPolicy()
	cache := cache.NewCache(config.CACHE_DEFAULT_SIZE, randomPolicy)
	cacheService := service.NewCacheService(cache)

	// Seed the cache with random data
	seeding.SeedWithRandomData(cache)

	http.HandleFunc("/get", cacheService.Get)
	http.HandleFunc("/set", cacheService.Set)
	http.HandleFunc("/delete", cacheService.Delete)
	http.HandleFunc("/clear", cacheService.Clear)

	fmt.Println("Starting cache server on :8080")
	http.ListenAndServe(":8080", nil)
}
