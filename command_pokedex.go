package main

import (
	"fmt"
)

// Add this function to your commands file or create a new file like commands/pokedex.go

func commandPokedex(cfg *config, args ...string) error {
	// Get the number of caught Pokemon
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon!")
		return nil
	}

	// Display the list of caught Pokemon
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
