package pokeapi

import (
	"net/http"
	"time"

	"github.com/saurabh0jha/pokedex/internal/pokecache"
)

// Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient
func NewClient(timeout, cacheInterval time.Duration) Client {
	pokeCache := pokecache.NewCache(
		cacheInterval,
	)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokeCache,
	}
}
