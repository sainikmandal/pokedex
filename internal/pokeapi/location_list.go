package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations retrieves locations from the PokeAPI
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RespShallowLocations{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var result RespShallowLocations
	err = json.Unmarshal(body, &result)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return result, nil
}
