package pokapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PokemonEncounters struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetPokemonEncounter(pUrl *string) (PokemonEncounters, error) {
	url := apiURL + "/location-area"
	if pUrl != nil {
		url = *pUrl
	}
	if val, ok := cache.Get(url); ok {
		PokemonE := PokemonEncounters{}
		err := json.Unmarshal(val, &PokemonE)
		if err != nil {
			return PokemonEncounters{}, err
		}
		return PokemonE, nil
	}
	// get the response if the there is no cache
	res, err := http.Get(url)
	if err != nil {
		return PokemonEncounters{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokemonEncounters{}, err
	}
	PokemonRes := PokemonEncounters{}
	err = json.Unmarshal(body, &PokemonRes)
	if err != nil {
		return PokemonEncounters{}, err
	}
	cache.Add(url, body)
	return PokemonRes, nil
}
