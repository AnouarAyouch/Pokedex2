package pokapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CatchPokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetPokemon(pokemonName string) (CatchPokemon, error) {
	url := apiURL + "/pokemon/" + pokemonName
	if val, ok := cache.Get(url); ok {
		PokemonE := CatchPokemon{}
		err := json.Unmarshal(val, &PokemonE)
		if err != nil {
			return CatchPokemon{}, err
		}
		return PokemonE, nil
	}
	// get the response if the there is no cache
	res, err := http.Get(url)
	if err != nil {
		return CatchPokemon{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return CatchPokemon{}, err
	}
	PokemonRes := CatchPokemon{}
	err = json.Unmarshal(body, &PokemonRes)
	if err != nil {
		return CatchPokemon{}, err
	}
	cache.Add(url, body)
	return PokemonRes, nil
}
