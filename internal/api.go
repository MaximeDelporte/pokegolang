package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MapResponse struct {
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Maps  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var currentResponse *MapResponse

func GetMaps(forward bool) {
	res, err := http.Get("https://pokeapi.co/api/v2/location")

	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	data := []byte(string(body))
	response := MapResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
    	fmt.Println(err)
	}
	fmt.Println(response.Next)
}
