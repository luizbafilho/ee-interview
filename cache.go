package main

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

// Provide a simple in-memory cache interface that can be use as an abstraction
// for any cache implementation. Allowing having a redis or memcached implementation if needed.
type cacher interface {
	Get(key string) (string, bool)
	Set(key string, value string) error
}

type cache struct {
	memCache *expirable.LRU[string, string]
}

func (c *cache) Get(key string) (string, bool) {
	return c.memCache.Get(key)
}

func (c *cache) Set(key string, value string) error {
	c.memCache.Add(key, value)
	return nil
}

func NewCache() cacher {
	// Picking arbitrary values for the cache size and expiration time
	// this should be adjusted based on the expected load and memory constraints
	lruCache := expirable.NewLRU[string, string](1000, nil, time.Second*30)

	return &cache{
		memCache: lruCache,
	}
}
