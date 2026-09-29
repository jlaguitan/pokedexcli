package main

import (
	"pokedexcli/internal/pokecache"
	"time"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		cache:    pokecache.NewCache(5 * time.Second),
	}

	startRepl(cfg)
}
