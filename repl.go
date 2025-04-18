package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userCommand := cleanInput(scanner.Text())[0]
			command, found := getCommands()[userCommand]
			if !found {
				fmt.Println("Unknown command")
			} else {
				err := command.callback()
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


type cliCommand struct {
	name string
	description string
	callback func() error
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
	}
}




 
