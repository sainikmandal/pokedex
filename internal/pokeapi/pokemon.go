// internal/pokeapi/pokemon.go
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Pokemon represents a Pokemon from the API
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

// GetPokemon retrieves information about a specific Pokemon by name
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)

	// Check if the response is in the cache
	if data, found := c.cache.Get(url); found {
		// Found in cache, decode and return
		fmt.Println("Cache hit:", url)
		var result Pokemon
		err := json.Unmarshal(data, &result)
		if err != nil {
			return Pokemon{}, err
		}
		return result, nil
	}

	// Not in cache, make the HTTP request
	fmt.Println("Cache miss:", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Add the response to the cache
	c.cache.Add(url, body)

	var result Pokemon
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Pokemon{}, err
	}

	return result, nil
}
