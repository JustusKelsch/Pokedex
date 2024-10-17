package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, pokemon []string) error {

	if len(pokemon) < 2 {
		return fmt.Errorf("no pokemon selected")
	}

	pokemonResp, err := config.pokeapiClient.GetSelectedPokemon(pokemon[1])
	if err != nil {
		return fmt.Errorf("invalid pokemon: %s", pokemon[1])
	}

	fmt.Printf("Throwing a pokeball at %s...\n", pokemonResp.Name)

	attempt := rand.Intn(pokemonResp.BaseExperience)

	if attempt > 40 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResp.Name)
	config.pokedex[pokemonResp.Name] = pokemonResp

	return nil
}