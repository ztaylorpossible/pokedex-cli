package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"github.com/ztaylorpossible/pokedex-cli/internal/pokeapi"
	"github.com/ztaylorpossible/pokedex-cli/internal/pokecache"
)

type config struct {
	commands  map[string]cliCommand
	next      string
	previous  string
	catchRoll int
	pokemon   map[string]pokeapi.PokemonInfo
	cache     pokecache.Cache
}

func main() {
	conf := config{
		commands: map[string]cliCommand{
			"help": {
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},
			"map": {
				name: "map",
				description: "List the next 20 locations",
				callback: commandMap,
			},
			"mapb": {
				name: "mapb",
				description: "List the previous 20 locations",
				callback: commandMapb,
			},
			"explore": {
				name: "explore",
				description: "Explore an area for pokemon (explore [location-name])",
				callback: commandExplore,
			},
			"catch": {
				name: "catch",
				description: "Attempt to catch a pokemon (catch [pokemon-name])",
				callback: commandCatch,
			},
			"exit": {
				name: "exit",
				description: "Exit the Pokedex",
				callback: commandExit,
			},
		},
		next: "",
		previous: "",
		catchRoll: 1000,
		pokemon: make(map[string]pokeapi.PokemonInfo),
		cache: pokecache.NewCache(5 * time.Minute),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) < 1 {
			fmt.Print("\n")
			continue
		}

		command, ok := conf.commands[words[0]]
		if !ok {
			fmt.Println("Unknown command. Use 'help' for a list of commands or 'exit' to quit.")
			continue
		}

		var params []string
		if len(words) > 1 {
			params = words[1:]
		}

		err := command.callback(&conf, params)
		if err != nil {
			fmt.Printf("Error processing command %s: %v\n", command.name, err)
		}
	}
}

