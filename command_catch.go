package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("usage: catch <pokemon>\n")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonDetails, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		// fmt.Printf("error --- %v", err)
		// fmt.Printf("pokemon --- %v", pokemonDetails)
		return err
	}

	chance := rand.Int()
	if chance < pokemonDetails.BaseExperience {
		fmt.Printf("%s escapted!\n", pokemonDetails.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonDetails.Name)

	cfg.pokedex[pokemonDetails.Name] = pokemonDetails

	return nil
}
