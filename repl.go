package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputText := cleanInput(scanner.Text())
		if len(inputText) == 0 {
			continue
		}

		fmt.Printf("Your command was: %v\n", inputText[0])
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
