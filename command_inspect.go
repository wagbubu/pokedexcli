package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := cfg.pokedex[pokemonName]; 
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("  - %s\n", val.Type.Name)
	}


	return nil
}