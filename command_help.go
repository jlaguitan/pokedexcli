package main

import (
	"fmt"
)

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
