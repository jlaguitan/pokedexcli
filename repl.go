package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands()map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if reader.Scan() {
			input := reader.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}
			
			commands := getCommands()
			if value, ok := commands[words[0]]; ok{
				err := value.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}			
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)

	commands := getCommands()
	for key, value := range(commands) {
		fmt.Println(key + ": " + value.description)
	}
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}