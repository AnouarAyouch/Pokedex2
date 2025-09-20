package pokapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AnouarAyouch/Pokedex2/internal/pokcache"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const apiURL = "https://pokeapi.co/api/v2"
const interval = 5 * time.Second

var cache = pokcache.NewCache(interval)

func ListLocations(pUrl *string) (RespShallowLocations, error) {
	url := apiURL + "/location-area"
	if pUrl != nil {
		url = *pUrl
	}
	if val, ok := cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}
	// get the response if the there is no cache
	res, err := http.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return RespShallowLocations{}, err
	}
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	cache.Add(url, body)
	return locationsResp, nil
}
