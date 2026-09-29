package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, s string) error {
	var url string
	if cfg.next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.next
	}

	var results pokeapi.LocationAreasResponse
	if val, ok := cfg.cache.Get(url); ok {
		err := json.Unmarshal(val, &results)
		if err != nil {
			return (err)
		}
	} else {
		res, err := http.Get(url) //response for GET
		if err != nil {
			return (err)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return (err)
		}
		defer res.Body.Close()

		cfg.cache.Add(url, data)

		err = json.Unmarshal(data, &results)
		if err != nil {
			return (err)
		}
	}

	cfg.next = results.Next         //check kung ano nangyayari
	cfg.previous = results.Previous //check kung ano nangyayari
	for _, result := range results.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(cfg *config, s string) error {
	//code mapb body, url = cfg.previous
	var url string
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}

	url = *cfg.previous

	var results pokeapi.LocationAreasResponse
	if val, ok := cfg.cache.Get(url); ok {
		err := json.Unmarshal(val, &results)
		if err != nil {
			return (err)
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body) //converts res.Body into []byte
		if err != nil {
			return err
		}

		cfg.cache.Add(url, data)

		err = json.Unmarshal(data, &results)
		if err != nil {
			return err
		}
	}

	cfg.previous = results.Previous
	cfg.next = results.Next
	for _, result := range results.Results {
		fmt.Println(result.Name)
	}

	return nil
}
