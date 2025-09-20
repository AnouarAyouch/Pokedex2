package main

import (
	"fmt"
	"strings"

	"github.com/AnouarAyouch/Pokedex2/internal/pokapi"
)

const apiURL = "https://pokeapi.co/api/v2"

func commandExplore(c *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location-name> ")
	}
	locationName := strings.Join(args, "")
	fullUrl := apiURL + "/location-area/" + locationName + "/"
	fmt.Println(fullUrl)
	area, err := pokapi.GetPokemonEncounter(&fullUrl)
	if err != nil {
		return fmt.Errorf("failed to get encounters: %w", err)
	}
	fmt.Printf("Location Area: %s\n", area.Name)
	fmt.Println("Pokémon Encounters:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}
