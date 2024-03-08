// https://pokeapi.co/api/v2/location-area/1
package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PokemonsResponse struct {
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

/*
	Public Interface
*/

func GetPokemonsFrom(city string) (*PokemonsResponse, error) {
	localResponse, ok := getLocalPokemons(city)
	if ok {
		return localResponse, nil
	}

	remoteResponse, err := getRemotePokemons(city)
	return remoteResponse, err
}

/*
	Private Interface
*/

func getLocalPokemons(city string) (*PokemonsResponse, bool) {
	cacheNotExists := true

	if cache != nil && cache.cacheEntry != nil {
		cacheNotExists = false
	}

	if cacheNotExists {
		cache = NewCache(10 * time.Second)
		return nil, false
	}

	key := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s-area", city)
	body, ok := cache.Get(key)
	if !ok {
		return nil, false
	}

	response, err := getPokemonResponseFrom(body)
	if err != nil {
		return nil, false
	}

	return response, true
}

// getRemotePokemons: Get REMOTE pokemons from the API.
func getRemotePokemons(city string) (*PokemonsResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s-area", city)
	fmt.Println("url:", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	if res.StatusCode > 299 && res.StatusCode < 400 {
		errString := fmt.Sprintf("Response failed.\nStatus code: %d\nBody: %s\n", res.StatusCode, body)
		err := errors.New(errString)
		return nil, err
	}

	if res.StatusCode > 400 {
		if res.StatusCode == 404 {
			return nil, errors.New("No pokemons found in this area.")
		} else {
			errString := fmt.Sprintf("Response failed.\nStatus code: %d\nBody: %s\n", res.StatusCode, body)
			err := errors.New(errString)
			return nil, err
		}
	}

	cache.Add(url, body)
	return getPokemonResponseFrom(body)
}

func getPokemonResponseFrom(byte []byte) (*PokemonsResponse, error) {
	response := PokemonsResponse{}

	err := json.Unmarshal(byte, &response)
	if err != nil {
		error := errors.New("A problem happened when the program tries to decode the response.")
		return nil, error
	}

	return &response, nil
}
