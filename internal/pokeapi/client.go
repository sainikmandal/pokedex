package pokeapi

import (
	"net/http"
	"time"
)

// Client is the PokeAPI client
type Client struct {
	httpClient http.Client
}

// New creates a new PokeAPI client
// NewClient creates a new PokeAPI client
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
