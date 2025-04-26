package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wagbubu/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			var mapName string
			var userCommand string
			userInput := cleanInput(scanner.Text())
			userCommand = cleanInput(scanner.Text())[0]
			if len(userInput) > 1 {
				mapName = cleanInput(scanner.Text())[1]
			}
			command, found := getCommands()[userCommand]
			if !found {
				fmt.Println("Unknown command")
			} else {
				err := command.callback(cfg, mapName)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
	}
}

func cleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	cleanStrArr := strings.Fields(lowerCased)
	return cleanStrArr
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},

		"help": {
			name: "help",
			description: "describe all commands",
			callback: commandHelp,
		},

		"map": {
			name: "map",
			description: "displays the names of first or next 20 location areas in the pokemon world.",
			callback: commandMapf,
		},

		"mapb": {
			name: "mapb",
			description: "displays the names of previous 20 location areas in the pokemon world.",
			callback: commandMapb,
		},

		"explore": {
			name: "explore <area-name>",
			description: "show list of pokemon encountered in an area",
			callback: commandExplore,
		},
	}
}




 
