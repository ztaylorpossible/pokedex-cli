package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/ztaylorpossible/pokedex-cli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := repl.CleanInput(text)
		fmt.Printf("Your command was: %s\n", words[0])
	}
}

