package main

import (
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"time"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		cache:    pokecache.NewCache(5 * time.Second),
		pokedex:  make(map[string]pokeapi.PokemonInfo),
	}

	startRepl(cfg)
}
