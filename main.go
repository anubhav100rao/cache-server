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
	fmt.Println("Using RandomEvictionPolicy")

	fmt.Println("Want to edit configs?")
	
	var cacheSize int
	fmt.Print("Enter cache size: ")
	fmt.Scan(&cacheSize)

	fmt.Println("Enter eviction policy: ")
	fmt.Println("1. RandomEvictionPolicy")
	fmt.Println("2. LRUEvictionPolicy")
	fmt.Println("3. LIFOEvictionPolicy")
	fmt.Println("4. FIFOEvictionPolicy")
	

	// Use a custom eviction policy, here we use RandomEvictionPolicy as an example
	randomPolicy := eviction.NewRandomEvictionPolicy()
	cache := cache.NewCache(config.CACHE_DEFAULT_SIZE, randomPolicy)
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
