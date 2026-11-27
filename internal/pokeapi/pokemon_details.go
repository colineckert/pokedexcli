package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	pokemon := Pokemon{}

	val, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, err

	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, err
	}

	return pokemon, nil
}
