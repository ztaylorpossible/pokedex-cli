package main

import (
	"fmt"
	"bufio"
	"os"
)

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
}

func main() {
	options := config{
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

		command, ok := options.commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&options)
		if err != nil {
			fmt.Printf("Error processing command %s: %v\n", command.name, err)
		}
	}
}

