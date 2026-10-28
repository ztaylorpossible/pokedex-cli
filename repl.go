package main

import (
	"strings"
	"os"
	"fmt"
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

func commandHelp(options *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, val := range options.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandExit(options *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
