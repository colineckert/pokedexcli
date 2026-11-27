package pokeapi

import (
	"encoding/json"
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
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return LocationAreasRes{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationsRes, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationsRes, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationsRes, nil
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return locationsRes, nil
	}

	return locationsRes, nil
}
