package pokeapi

import (
	"net/http"
	"encoding/json"
)

type ResourceResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResourceList struct {
	Count    int              `json:count`
	Next     string           `json:next`
	Previous string           `json:previous`
	Results  []ResourceResult `json:results`
}

func GetLocations(url string) (ResourceList, error) {
	client := &http.Client{}
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResourceList{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return ResourceList{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var locations ResourceList
	err = decoder.Decode(&locations)
	if err != nil {
		return ResourceList{}, err
	}

	return locations, nil
}
