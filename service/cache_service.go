package service

import (
	"encoding/json"
	"net/http"

	"github.com/anubhav100rao/cache_server/cache"
)

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService(cache *cache.Cache) *CacheService {
	return &CacheService{cache: cache}
}

func (cs *CacheService) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("key")
	if key == "" {
		responses := []map[string]interface{}{}
		for k, v := range cs.cache.GetAll() {
			responses = append(responses, map[string]interface{}{"key": k, "value": v})
		}

		json.NewEncoder(w).Encode(responses)
		return
	}

	value, exists := cs.cache.Get(key)
	if !exists {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": value})
}

func (cs *CacheService) Set(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs.cache.Set(requestData.Key, requestData.Value)
	w.WriteHeader(http.StatusNoContent)
}

func (cs *CacheService) Delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
		return
	}

	cs.cache.Delete(key)
	w.WriteHeader(http.StatusNoContent)
}

func (cs *CacheService) Clear(w http.ResponseWriter, r *http.Request) {
	cs.cache.Clear()
	w.WriteHeader(http.StatusNoContent)
}
