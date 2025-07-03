package main

import (
	"fmt"
	"github.com/SGWinter/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	locations, err := pokeapi.LocationList()
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}
