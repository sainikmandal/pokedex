package main

import (
	"math/rand"
	"time"

	"github.com/sainikmandal/pokedex/internal/pokeapi"
)

func init() {
	// For Go 1.20+ (the modern approach)
	// This replaces the deprecated rand.Seed method
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
