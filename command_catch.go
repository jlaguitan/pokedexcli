package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, p string) error {
	if p == "" {
		return errors.New("no pokemon input")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + p

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("pokemon not found: %s", p)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var results pokeapi.PokemonInfo
	err = json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + p + "...")
	roll := rand.Intn(results.BaseExperience)
	threshold := 40

	if roll < threshold {
		fmt.Println(p + " was caught!")
		fmt.Println("You may now inspect it with the inspect command")
		cfg.pokedex[p] = results
	} else {
		fmt.Println(p + " escaped!")
	}

	return nil
}

//url := "https://pokeapi.co/api/v2/pokemon/" + //id or name
//takes the name of a pokemon as an argument
//Pokedex > catch pickachu
//Throwing a Pokeball at pikachu...
//pickachu escaped!
//Pokedex > catch pikachu
//Throwing a Pokeball at pikachu...
//pikachu was caught!

//use Pokemon endpoint to get info about pokemon by name

//give user a chance to catch using math/rand package

//pokemon's base experience > basis for catching
// the higher, the harder to catch

//one pokemon is caught, add it to user's pokedex.
//map[string]Pokemon
