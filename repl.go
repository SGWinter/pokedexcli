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

		commandName := inputText[0]

		command, exists := getCommand()[commandName]

		if exists {
			err := command.callbak()
			if err != nil {
				fmt.Println(err)
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callbak     func() error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callbak:     commandExit,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
