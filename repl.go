package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/colineckert/pokedexcli/internal/pokeapi"
)

// shared state
type config struct {
	commands             map[string]cliCommand
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    handler
}

type handler func(cfg *config, args []string) error

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List next 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore location are and list pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		cleanedInput := cleanInput(input)
		commandName := cleanedInput[0]
		supportedCommands := getCommands()
		command, ok := supportedCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		// Get command args
		args := cleanedInput[1:]

		command.callback(cfg, args)
	}
}
