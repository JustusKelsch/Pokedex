package main

import "fmt"

func commandExplore(config *config, args []string) error {

	if len(args) < 2 {
		return fmt.Errorf("no location specified")
	}

	pokemonResp, err := config.pokeapiClient.ListPokemon(args[1])
	if err != nil {
		return fmt.Errorf("invalid location")
	}

	fmt.Printf("Exploring %s...", args[1])
	fmt.Println()

	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %s", encounter.Pokemon.Name)
		fmt.Println()
	}

	return nil
}