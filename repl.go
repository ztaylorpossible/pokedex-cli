package main

import (
	"strings"
	"os"
	"fmt"
	"github.com/ztaylorpossible/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, params []string) error
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	split := strings.Split(text, " ")
	var result []string
	for _, s := range split {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

func commandHelp(conf *config, params []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, val := range conf.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandMap(conf *config, params []string) error {
	result, err := pokeapi.GetLocations(conf.next, &conf.cache)
	if err != nil {
		return err
	}
	conf.next = result.Next
	conf.previous = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(conf *config, params []string) error {
	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	result, err := pokeapi.GetLocations(conf.previous, &conf.cache)
	if err != nil {
		return err
	}
	conf.next = result.Next
	conf.previous = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(conf *config, params []string) error {
	if params == nil || len(params) < 1 {
		return fmt.Errorf("No location provided to explore (Usage: explore [location-name])")
	}
	location := params[0]

	fmt.Printf("Exploring %s...\n", location)
	areaInfo, err := pokeapi.GetLocationArea(location, &conf.cache)
	if err != nil {
		return err
	}

	if len(areaInfo.Pokemon_Encounters) < 1 {
		fmt.Println("No Pokemon found")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range areaInfo.Pokemon_Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandExit(conf *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
