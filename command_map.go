package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"
	"pokedexcli/internal/pokeapi"
	"errors"
)



func commandMap(cfg *config) error {
	var url string
	if cfg.next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.next
	}

	res, err := http.Get(url)	//response for GET
	if err != nil {
		return(err)
	} 
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return(err)
	}
	defer res.Body.Close()

	var results pokeapi.LocationAreasResponse

	err = json.Unmarshal(data, &results)
	if err != nil {
		return(err)
	}

	cfg.next = results.Next			//check kung ano nangyayari
	cfg.previous = results.Previous	//check kung ano nangyayari
	for _, result := range results.Results {
		fmt.Println(result.Name)
	}
	return nil

}

func commandMapb(cfg *config) error {
	//code mapb body, url = cfg.previous
	var url string
	if cfg.previous == nil {
		return errors.New("you're on the first page")	
	} 

	url = *cfg.previous
	
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body) //converts res.Body into []byte
	if err != nil {
		return err
	}


	var results pokeapi.LocationAreasResponse
	err = json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	cfg.previous = results.Previous
	cfg.next = results.Next
	for _,result := range results.Results {
		fmt.Println(result.Name)
	}
	
	return nil
}