package main

import (
	"fmt"
	"github.com/SGWinter/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	o := 20
	locations, err := pokeapi.LocationList(o)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func commandMapb() error {
	o := -20
	locations, err := pokeapi.LocationList(o)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}
