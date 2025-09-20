package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, args []string) error
}

// it runs the command

// Exit command
func commandExit(c *Config, args []string) error {
	fmt.Println("")
	fmt.Println("Closing the Pokedex... Goodbye! ")
	os.Exit(0)
	return nil
}

// Help command
func commandHelp(c *Config, args []string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range Commands() {
		fmt.Printf(" %s : %s\n", cmd.name, cmd.description)
	}
	return nil
}

// func for all the commands , then it return the chosen cmd
func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get the pokemen encounter in the provided area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch the pokemon ",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "it shows you all the pokemon that you caught ",
			callback:    showAllPokemonInPokedex,
		},
	}
}
