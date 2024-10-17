package main

import "fmt"

func commandInspect(config *config, pokemon []string) error {
	if len(pokemon) < 2 {
		return fmt.Errorf("no pokemon selected")
	}

	pokemonResp, ok := config.pokedex[pokemon[1]]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemonResp.Name)
	fmt.Printf("Height: %v\n", pokemonResp.Height)
	fmt.Printf("Weight: %v\n", pokemonResp.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemonResp.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types: ")
	for _, types := range pokemonResp.Types {
		fmt.Printf("  - %v\n", types.Type.Name)
	}

	return nil
}