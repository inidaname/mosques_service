package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheService handles the in-memory caching
type CacheService struct {
	cache *cache.Cache
}

// NewCacheService initializes a new cache with default expiration times
func NewCacheService(defaultExpiration, cleanupInterval time.Duration) *CacheService {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &CacheService{cache: c}
}

// Set adds a value to the cache with a specified key
func (c *CacheService) Set(key string, value interface{}, ttl time.Duration) {
	c.cache.Set(key, value, ttl)
}

// Get retrieves a value from the cache by key
func (c *CacheService) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

// Delete removes a value from the cache by key
func (c *CacheService) Delete(key string) {
	c.cache.Delete(key)
}

// Clear clears all the values in the cache
func (c *CacheService) Clear() {
	c.cache.Flush()
}

// SetDefault sets a default value in the cache if the key does not exist
func (c *CacheService) SetDefault(key string, value interface{}) {
	// Set the value only if it does not already exist in the cache
	c.cache.SetDefault(key, value)
}
