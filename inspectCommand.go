package main

import "fmt"

func commandInspect(c *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon-name> ")
	}
	nameP := args[0]
	if _, exists := caughtPokemon[nameP]; exists {
		for _, p := range caughtPokemon {
			fmt.Printf("Name: %s\n", p.Name)
			fmt.Printf("Height: %d\n", p.Height)
			fmt.Printf("Weight: %d\n", p.Weight)
			for _, s := range p.Stats {
				fmt.Printf(" - %s: %d\n", s.Stat.Name, s.BaseStat)
			}
			fmt.Println("Types:")
			for _, s := range p.Types {
				fmt.Printf("- %s\n", s.Type.Name)
			}
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
