package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, s string) error {
	var url string
	url = "https://pokeapi.co/api/v2/location-area/" + s
	var results pokeapi.LocationAreaDetail

	if value, ok := cfg.cache.Get(url); ok {
		err := json.Unmarshal(value, &results)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, data)

		err = json.Unmarshal(data, &results)
		if err != nil {
			return err
		}
	}

	fmt.Println("Exploring " + s + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range results.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}
