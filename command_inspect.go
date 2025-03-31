// command_inspect.go

package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	// Check if the user has caught this Pokemon
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Display Pokemon information
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	// Display stats
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	// Display types
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
