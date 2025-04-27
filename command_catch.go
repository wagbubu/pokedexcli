package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	maxExp = 635
)

func commandCatch(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("please follow the command format: catch <pokemon-name>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	chanceOfCatching := 1 - float64(pokemonInfo.BaseExperience)/float64(maxExp)
	randomRoll := rand.Float64()
	
	if randomRoll < chanceOfCatching {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}