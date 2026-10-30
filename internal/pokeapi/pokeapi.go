package pokeapi

import (
	"fmt"
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

type LocationArea struct {
	Id                     int                   `json:id`
	Name                   string                `json:name`
	Game_Index             int                   `json:game_index`
	Encounter_Method_Rates []EncounterMethodRate `json:encounter_method_rates`
	Location               ResourceResult        `json:location`
	Names                  []NameResult          `json:names`
	Pokemon_Encounters     []PokemonEncounter    `json:pokemon_encounters`
}

type EncounterMethodRate struct {
	Encounter_Method ResourceResult            `json:encounter_method`
	Version_Details  []EncounterVersionDetails `json:version_details`
}

type EncounterVersionDetails struct {
	Rate    int            `json:rate`
	Version ResourceResult `json:version`
}

type NameResult struct {
	Name     string         `json:name`
	Language ResourceResult `json:language`
}

type PokemonEncounter struct {
	Pokemon         ResourceResult           `json:pokemon`
	Version_Details []VersionEncounterDetail `json:version_details`
}

type VersionEncounterDetail struct {
	Version           ResourceResult `json:version`
	Max_Chance        int            `json:max_chance`
	Encounter_Details []Encounter    `json:encounter_details`
}

type Encounter struct {
	Min_Level        int              `json:min_level`
	Max_Level        int              `json:max_level`
	Condition_Values []ResourceResult `json:condition_values`
	Chance           int              `json:chance`
	Method           ResourceResult   `json:method`
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

func GetLocationArea(area string, cache *pokecache.Cache) (LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	data, ok := cache.Get(url)
	if !ok {
		client := &http.Client{}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		res, err := client.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}

		var locationArea LocationArea
		err = json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		cache.Add(url, data)

		return locationArea, nil
	}

	var locationArea LocationArea
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
