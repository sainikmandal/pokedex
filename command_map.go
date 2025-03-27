package main

import (
	"errors"
	"fmt"
)

// Command: Show the map of locations
func commandMap(cfg *config, args ...string) error {
	// Get the locations using the client
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Update pagination URLs for future navigation
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	// Print the location names
	fmt.Println("Locations:")
	for _, loc := range resp.Results {
		fmt.Println("-", loc.Name)
	}

	return nil
}

// Command: Show the previous page of locations
func commandMapb(cfg *config, args ...string) error {
	// Check if we're already at the first page
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Get the previous page of locations
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Update pagination URLs for future navigation
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	// Print the location names
	fmt.Println("Locations (previous page):")
	for _, loc := range resp.Results {
		fmt.Println("-", loc.Name)
	}

	return nil
}
