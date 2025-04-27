package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, _ string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.pokedex) == 0 {
		return errors.New("(empty)")
	}
	for pokemonName, _ := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemonName)
	}
	return nil
}