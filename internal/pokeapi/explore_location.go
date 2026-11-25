package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + location

	locationDetails := LocationDetails{}

	val, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &locationDetails)
		if err != nil {
			return LocationDetails{}, err
		}
		return locationDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationDetails)
	if err != nil {
		return LocationDetails{}, err
	}

	return locationDetails, nil
}
