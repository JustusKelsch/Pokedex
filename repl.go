package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/justuskelsch/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex map[string]pokeapi.Pokemon
}

func startRepl(config *config) {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(config, words)
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
	name string
	description string
	callback func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"inspect": {
			name: "inspect",
			description: "View the stats of a pokemon you have caught",
			callback: commandInspect,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a specified pokemon",
			callback: commandCatch,
		},
		"explore": {
			name: "explore",
			description: "Lists the pokemon found in the specified location",
			callback: commandExplore,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations",
			callback: commandMapB,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}