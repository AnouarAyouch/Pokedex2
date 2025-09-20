package main

import (
	"fmt"
	"math/rand"

	"github.com/AnouarAyouch/Pokedex2/internal/pokapi"
)

//	type Pokemon struct {
//		Name           string
//		BaseExperience int
//	}
var caughtPokemon = map[string]pokapi.CatchPokemon{}

func commandCatch(c *Config, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon-name> ")
	}
	pokemonName := args[0]
	pokemon, err := pokapi.GetPokemon(pokemonName)

	if err != nil {
		return fmt.Errorf("failed to get thr pokemon info: %w", err)
	}
	fmt.Printf("Throwing a Pokeball at %s... \n", pokemon.Name)
	baseExperience := pokemon.BaseExperience
	maxBase := 310
	catchChance := (maxBase - baseExperience) * 100 / maxBase
	if catchChance < 0 {
		catchChance = 0
	}
	//fmt.Println("Catch chance:", catchChance, "%")
	random := rand.Intn(100)
	if random < catchChance {
		fmt.Printf("%s was caught! \n", pokemon.Name)
		caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s  escaped!\n", pokemon.Name)

	}
	return nil
}

func showAllPokemonInPokedex(c *Config, args []string) error {

	fmt.Println("=== Your Pokedex ===")
	for _, p := range caughtPokemon {
		fmt.Printf("- %s (Base Exp: %d)\n", p.Name, p.BaseExperience)
	}
	return nil
}
