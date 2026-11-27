package main

import (
	"time"

	"github.com/colineckert/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
