package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, mapName string) error {
	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _ , loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, mapName string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}	

	locationRes, err :=  cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationRes.Next
	cfg.prevLocationsURL = locationRes.Previous

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}