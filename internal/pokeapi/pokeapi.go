package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
	"github.com/ztaylorpossible/pokedex-cli/internal/pokecache"
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

func GetLocations(url string, cache *pokecache.Cache) (ResourceList, error) {
	data, ok := cache.Get(url)
	if !ok {
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

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return ResourceList{}, err
		}

		var locations ResourceList
		err = json.Unmarshal(data, &locations)
		if err != nil {
			return ResourceList{}, err
		}

		cache.Add(url, data)

		return locations, nil
	}

	var locations ResourceList
	err := json.Unmarshal(data, &locations)

	if err != nil {
		return ResourceList{}, err
	}

	return locations, nil
}
