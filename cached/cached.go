package cached

import (
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/go-redis/redis/v7"
)

// Manager Manager
var Manager *cache.ChainCache

// Init Init
func Init() error {
	// Initialize Ristretto cache and Redis client
	_ristrettoCache, err := ristretto.NewCache(&ristretto.Config{NumCounters: 1000, MaxCost: 100, BufferItems: 64})
	if err != nil {
		panic(err)
	}

	_redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	// Initialize stores
	ristrettoStore := store.NewRistretto(_ristrettoCache, nil)
	redisStore := store.NewRedis(_redisClient, &store.Options{})

	// Initialize chained cache
	Manager = cache.NewChain(
		cache.New(ristrettoStore),
		cache.New(redisStore),
	)
	return nil
}
