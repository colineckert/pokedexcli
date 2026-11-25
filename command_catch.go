package main

import "fmt"

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("usage: catch <pokemon>\n")
	}

	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...", pokemon)

	return nil
}
