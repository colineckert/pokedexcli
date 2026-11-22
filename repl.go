package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input := ""
		if scanner.Scan() {
			input = input + scanner.Text()
		}

		if input == "exit" {
			fmt.Println("Exiting Pokedex REPL.")
			break
		}

		words := cleanInput(input)

		fmt.Printf("Your command was: %s", words[0])
		fmt.Println()
	}
}
