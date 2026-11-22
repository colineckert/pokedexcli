package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	// state
	commands map[string]cliCommand
}

func getConfig() config {
	return config{
		commands: getCommands(),
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

		command.callback(cfg)
	}
}
