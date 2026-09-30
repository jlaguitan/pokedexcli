package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, p string) error {
	if p == "" {
		return errors.New("no pokemon input")
	}

	pokeInfo, ok := cfg.pokedex[p]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name: " + p)
		fmt.Printf("Height: %v\n", pokeInfo.Height)
		fmt.Printf("Weight: %v\n", pokeInfo.Weight)
		fmt.Println("Stats:")

		for _, stat := range pokeInfo.Stats {
			fmt.Printf("   -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")

		for _, typez := range pokeInfo.Types {
			fmt.Printf("   -%v\n", typez.Type.Name)
		}
	}
	return nil
}

//Pokedex > inspect pidgey
//you have not caught that pokemon

//Pokedex > inspect pidgey
//Name: pidgey
//Height: 3
//Weight: 18
//Stats:
//	-hp: 40
//	-attack: 45
//	-defense: 40
//	-special-attack: 35
//	-special-defense: 35
//	-speed: 56
//Types:
//	- normal
//	- flying
