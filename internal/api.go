package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func GetNextMaps() (*MapResponse, error) {
	return getMaps(true)
}

func GetPreviousMaps() (*MapResponse, error) {
	return getMaps(false)
}

func getMaps(forward bool) (*MapResponse, error) {
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

	data := []byte(body)
	response := MapResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		error := errors.New("A problem happened when the program tries to decode the response.")
		return nil, error
	}

	currentResponse = &response
	return &response, nil
}

func getURLFromCurrentResponse(forward bool) (*string, error) {
	if currentResponse == nil && forward == false {
		return nil, errors.New("You can't go backward !")
	}

	url := "https://pokeapi.co/api/v2/location"

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
