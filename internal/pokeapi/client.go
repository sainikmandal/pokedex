package pokeapi

import (
	"net/http"
	"time"

	"github.com/sainikmandal/pokedex/internal/pokecache"
)

// Client is the PokeAPI client
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient creates a new PokeAPI client with caching
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(5 * time.Minute), // Cache entries for 5 minutes
	}
}
