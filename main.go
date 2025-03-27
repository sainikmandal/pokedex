package main

import (
	"github.com/sainikmandal/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(cfg)
}
