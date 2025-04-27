package main

import (
	"time"

	"github.com/wagbubu/pokedexcli/internal/pokeapi"
)

func main () {
	pokemonApiClient := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	cfg := &config{pokeapiClient: pokemonApiClient, pokedex: make(map[string]pokeapi.Pokemon)} 
	startRepl(cfg)
}

