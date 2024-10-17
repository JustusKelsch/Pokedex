package main

import "fmt"

func commandPokedex(config *config, args []string) error {

	if len(config.pokedex) == 0 {
		fmt.Println("You have not caught any pokemon.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Println("  -", pokemon.Name)
	}

	return nil
}