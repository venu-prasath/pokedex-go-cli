package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/venu-prasath/pokedexcli/internal/pokeapi"
)

type config struct {
	nextLocationUrl     *string
	previousLocationUrl *string
	pokeapiClient       pokeapi.Client
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Incorrect input. Type help for help")
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	callback    func(*config) error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the app",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get next set of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous set of locations",
			callback:    commandMapb,
		},
	}
}
