package main

import (
	"fmt"
	"github.com/SGWinter/pokedexcli/internal"
)

func commandMap() error {
	locations := internal.locationListAPI()

	for location := range locations {
		fmt.Printf("%v\n", location.areas.name)
	}
	return nil
}
