package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"github.com/ztaylorpossible/pokedex-cli/internal/pokecache"
)

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
	cache    pokecache.Cache
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
			"exit": {
				name: "exit",
				description: "Exit the Pokedex",
				callback: commandExit,
			},
		},
		next: "",
		previous: "",
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
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&conf)
		if err != nil {
			fmt.Printf("Error processing command %s: %v\n", command.name, err)
		}
	}
}

