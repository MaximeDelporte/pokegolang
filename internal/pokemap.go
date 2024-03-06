package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MapResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Maps     []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

/*
	Public Interface
*/

func GetNextMaps() (*MapResponse, error) {
	return getMaps(true)
}

func GetPreviousMaps() (*MapResponse, error) {
	return getMaps(false)
}

/*
	Private Interface
*/

var currentResponse *MapResponse
var cache *Cache

func getMaps(forward bool) (*MapResponse, error) {
	localResponse, ok := getLocalMaps(forward)
	if ok {
		return localResponse, nil
	}

	remoteResponse, err := getRemoteMaps(forward)
	return remoteResponse, err
}

// getLocalMaps: Get LOCAL maps from the cache.
func getLocalMaps(forward bool) (*MapResponse, bool) {
	if cache == nil {
		cache = NewCache(10 * time.Second)
		return nil, false
	}

	url, err := getURLFromCurrentResponse(forward)
	if err != nil {
		return nil, false
	}

	body, ok := cache.Get(*url)
	if !ok {
		return nil, false
	}

	response, err := getMapResponseFrom(body)
	if err != nil {
		return nil, false
	}

	return response, true
}

// getRemoteMaps: Get REMOTE maps from the API.
func getRemoteMaps(forward bool) (*MapResponse, error) {
	url, err := getURLFromCurrentResponse(forward)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()
	if res.StatusCode > 299 {
		errString := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		err := errors.New(errString)
		return nil, err
	}

	cache.Add(*url, body)
	return getMapResponseFrom(body)
}

/*
	Common Methods
*/

func getURLFromCurrentResponse(wantToMoveForward bool) (*string, error) {
	currentResponseExists := currentResponse != nil

	if currentResponseExists {
		if wantToMoveForward {
			if currentResponse.Next != nil {
				return currentResponse.Next, nil
			} else {
				return nil, errors.New("All the maps has been discovered.")
			}
		} else {
			if currentResponse.Previous != nil {
				return currentResponse.Previous, nil
			} else {
				return nil, errors.New("You can't go backward !")
			}
		}
	} else {
		if wantToMoveForward {
			url := "https://pokeapi.co/api/v2/location?offset=0&limit=20"
			return &url, nil
		} else {
			return nil, errors.New("You can't go backward !")
		}
	}
}

func getMapResponseFrom(byte []byte) (*MapResponse, error) {
	response := MapResponse{}

	err := json.Unmarshal(byte, &response)
	if err != nil {
		error := errors.New("A problem happened when the program tries to decode the response.")
		return nil, error
	}

	currentResponse = &response
	return &response, nil
}
