package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasRes, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationsRes := LocationAreasRes{}

	val, ok := c.cache.Get(url)
	if ok {
		fmt.Printf("Cache hit! url: %s, value: %s\n", url, string(val))
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return LocationAreasRes{}, err
		}
		return locationsRes, nil
	}

	fmt.Printf("Cache miss! url: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return LocationAreasRes{}, err
	}

	return locationsRes, nil
}
