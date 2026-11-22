package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasRes struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchLocationAreas(url string) (LocationAreasRes, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasRes{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreasRes{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return LocationAreasRes{}, err
	}

	var locAreaRes LocationAreasRes
	if err = json.Unmarshal(body, &locAreaRes); err != nil {
		return LocationAreasRes{}, fmt.Errorf("Failed to unmarshal response body: %s\nError: %v", body, err)
	}

	return locAreaRes, nil
}

func commandMap(cfg *config) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	locAreaRes, err := fetchLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Next = locAreaRes.Next
	cfg.Previous = locAreaRes.Previous

	for _, loc := range locAreaRes.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}
