package main

import (
	"fmt"
)

func commandMapB(cfg *config) error {
	url := cfg.Previous
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
