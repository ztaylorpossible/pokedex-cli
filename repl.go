package main

import (
	"strings"
	"os"
	"fmt"
	"github.com/ztaylorpossible/pokedex-cli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, val := range conf.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandMap(conf *config) error {
	result, err := pokeapi.GetLocations(conf.next)
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

func commandMapb(conf *config) error {
	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	result, err := pokeapi.GetLocations(conf.previous)
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

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
