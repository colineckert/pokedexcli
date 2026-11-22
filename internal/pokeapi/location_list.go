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

	locationsRes := LocationAreasRes{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return LocationAreasRes{}, err
	}

	return locationsRes, err
}
