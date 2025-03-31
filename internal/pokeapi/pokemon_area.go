// internal/pokeapi/pokemon_area.go
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreaResponse represents the response from the location-area endpoint
type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// GetLocationArea retrieves information about a specific location area
func (c *Client) GetLocationArea(areaName string) (LocationAreaResponse, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, areaName)

	// Check if the response is in the cache
	if data, found := c.cache.Get(url); found {
		// Found in cache, decode and return
		fmt.Println("Cache hit:", url)
		var result LocationAreaResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return result, nil
	}

	// Not in cache, make the HTTP request
	fmt.Println("Cache miss:", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Add the response to the cache
	c.cache.Add(url, body)

	var result LocationAreaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return result, nil
}

// GetPokemonInArea returns a list of Pokémon found in the specified location area
func (c *Client) GetPokemonInArea(areaName string) ([]string, error) {
	areaInfo, err := c.GetLocationArea(areaName)
	if err != nil {
		return nil, err
	}

	pokemonNames := make([]string, 0, len(areaInfo.PokemonEncounters))
	for _, encounter := range areaInfo.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	return pokemonNames, nil
}
