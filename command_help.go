package main

import (
	"fmt"
)


func commandHelp() error {
	commandList := getCommand()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for command := range commandList {
		fmt.Printf("%v: %v\n", commandList[command].name, commandList[command].description)
	}
	return nil
}
