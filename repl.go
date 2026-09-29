package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	commands map[string]cliCommand
	next     *string
	previous *string
	cache    *pokecache.Cache
	pokedex  map[string]pokeapi.PokemonInfo
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokemon list in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon and save it to the Pokedex",
			callback:    commandCatch,
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
			if value, ok := commands[words[0]]; ok {
				if len(words) >= 2 {
					err := value.callback(cfg, words[1])
					if err != nil {
						fmt.Println(err)
					}
				} else {
					err := value.callback(cfg, "")
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit(cfg *config, s string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, s string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)

	commands := cfg.commands
	for key, value := range commands {
		fmt.Println(key + ": " + value.description)
	}
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
