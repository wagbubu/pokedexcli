package main

import "fmt"

func commandHelp(cfg *config, mapName string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Printf("%s: %s \n",key, value.description)
	}
	fmt.Println()
	return nil
}