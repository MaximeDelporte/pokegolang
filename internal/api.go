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

var currentResponse *MapResponse
var cache *Cache

func GetNextMaps() (*MapResponse, error) {
	response, ok := getMapsFromCache(true)
	if !ok {
		return getMapsFromApi(true)
	}
	return response, nil
}

func GetPreviousMaps() (*MapResponse, error) {
	response, ok := getMapsFromCache(false)
	if !ok {
		return getMapsFromApi(false)
	}
	return response, nil
}

func getMapsFromCache(forward bool) (*MapResponse, bool) {
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

func getMapsFromApi(forward bool) (*MapResponse, error) {
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

	response, err := getMapResponseFrom(body)
	cache.Add(*url, body)
	return response, err
}

func getURLFromCurrentResponse(forward bool) (*string, error) {
	if currentResponse == nil && forward == false {
		return nil, errors.New("You can't go backward !")
	}

	url := "https://pokeapi.co/api/v2/location?offset=0&limit=20"

	if currentResponse != nil {
		if forward {
			if currentResponse.Next == nil {
				return nil, errors.New("All the maps has been discovered.")
			} else {
				return currentResponse.Next, nil
			}
		} else {
			if currentResponse.Previous == nil {
				return nil, errors.New("You can't go backward !")
			} else {
				return currentResponse.Previous, nil
			}
		}
	}

	return &url, nil
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
