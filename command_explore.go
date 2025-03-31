// commands/explore.go
package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]
	fmt.Printf("Exploring %s...\n", locationAreaName)

	pokemonNames, err := cfg.pokeapiClient.GetPokemonInArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, name := range pokemonNames {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
