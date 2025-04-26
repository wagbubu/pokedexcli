package main

import "fmt"

func commandExplore(cfg *config, mapName string) error {
	fmt.Printf("Exploring %v...\n", mapName)
	fmt.Println("Found Pokemon:")
	pokemons, err := cfg.pokeapiClient.PokemonList(mapName);
	if err != nil {
		return err
	}
	for _, v := range pokemons.PokemonEncounters {
		fmt.Printf("  - %s\n", v.Pokemon.Name)
	}
	return nil
}