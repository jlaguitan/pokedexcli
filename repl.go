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
	callback func(*config) error
}

type config struct {
	commands map[string]cliCommand
	next *string
	previous *string
}

func getCommands() map[string]cliCommand {
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
		"map": {
			name: "map",
			description: "Displays 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "map",
			description: "Displays previous 20 location areas",
			callback: commandMapb, 
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if reader.Scan() {
			input := reader.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}
			
			commands := cfg.commands
			if value, ok := commands[words[0]]; ok{
				err := value.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}			
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)

	commands := cfg.commands
	for key, value := range(commands) {
		fmt.Println(key + ": " + value.description)
	}
	return nil
}



func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}