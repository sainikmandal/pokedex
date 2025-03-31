// commands/catch.go
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Get Pokemon information
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Calculate catch chance - higher base experience means lower chance
	catchChance := calculateCatchChance(pokemon.BaseExperience)

	// Random number between 0 and 1
	randNum := rand.Float64()

	if randNum <= catchChance {
		// Pokemon was caught
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		// Pokemon escaped
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

// calculateCatchChance returns a catch probability based on base experience
// Lower base experience = higher catch chance
func calculateCatchChance(baseExp int) float64 {
	// Convert baseExp to a probability between 0.1 and 0.9
	// Higher baseExp = lower probability
	if baseExp <= 0 {
		return 0.9 // High chance for unknown/zero base experience
	}

	// Cap the max difficulty
	maxBaseExp := 600.0

	if float64(baseExp) > maxBaseExp {
		baseExp = int(maxBaseExp)
	}

	// Calculate probability (inverse relationship with baseExp)
	// This formula gives:
	// - Low baseExp (around 50) = ~0.8 chance
	// - Medium baseExp (around 150) = ~0.5 chance
	// - High baseExp (around 300+) = ~0.2 chance
	return 0.9 - 0.7*(float64(baseExp)/maxBaseExp)
}
