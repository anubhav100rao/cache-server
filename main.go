package main

import (
	"fmt"
	"time"

	"net/http"

	cache "github.com/anubhav100rao/cache_server/cache"
	"github.com/anubhav100rao/cache_server/config"
	eviction "github.com/anubhav100rao/cache_server/eviction"
	seeding "github.com/anubhav100rao/cache_server/seeding"
	service "github.com/anubhav100rao/cache_server/service"
)

func main() {

	fmt.Println("Starting cache server")
	fmt.Println("Cache size: ", config.CACHE_DEFAULT_SIZE)
	fmt.Println("Cache expiry: ", config.DURATION)
	fmt.Println("Using RandomEvictionPolicy By default")

	fmt.Println("Enter eviction policy: ")
	fmt.Println("1. RandomEvictionPolicy: ")
	fmt.Println("2. LRUEvictionPolicy: ")
	fmt.Println("3. LIFOEvictionPolicy: ")
	fmt.Println("4. FIFOEvictionPolicy: ")
	fmt.Println("5. LFUEvictionPolicy: ")

	var evictionPolicy int
	fmt.Scan(&evictionPolicy)

	var policy eviction.EvictionPolicy
	switch evictionPolicy {
	case 1:
		policy = eviction.NewRandomEvictionPolicy()
		fmt.Println("Using RandomEvictionPolicy")
	case 2:
		policy = eviction.NewLRUCachePolicy()
		fmt.Println("Using LRUEvictionPolicy")
	case 3:
		policy = eviction.NewLIFOQueuePolicy()
		fmt.Println("Using LIFOEvictionPolicy")
	case 4:
		policy = eviction.NewFIFOCachePolicy()
		fmt.Println("Using FIFOEvictionPolicy")
	case 5:
		policy = eviction.NewLFUCachePolicy()
		fmt.Println("Using LFUEvictionPolicy")
	default:
		policy = eviction.NewRandomEvictionPolicy()
		fmt.Println("Using RandomEvictionPolicy")
	}

	cache := cache.NewCache(config.CACHE_DEFAULT_SIZE, policy)
	cacheService := service.NewCacheService(cache)

	// cleaning up expired items every minute
	cache.StartCleanup(1 * time.Minute) // Set cleanup interval to 1 minute

	// Seed the cache with random data
	seeding.SeedWithRandomData(cache)

	http.HandleFunc("/get", cacheService.Get)
	http.HandleFunc("/set", cacheService.Set)
	http.HandleFunc("/delete", cacheService.Delete)
	http.HandleFunc("/clear", cacheService.Clear)

	fmt.Println("Starting cache server on :8080")
	http.ListenAndServe(":8080", nil)
}
