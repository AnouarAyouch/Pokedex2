package main

import (
	"fmt"

	"github.com/AnouarAyouch/Pokedex2/internal/pokapi"
)

type Config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

func commandMapf(c *Config, args []string) error {
	locationsResp, err := pokapi.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
func commandMapb(c *Config, args []string) error {
	locationsResp, err := pokapi.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
