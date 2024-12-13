package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type configuration struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *configuration
}

var commands map[string]cliCommand

func main() {

	commands = map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    CommandMap,
		},
		// "mapb": {
		// 	name:        "mapback",
		// 	description: "Lists previous location areas",
		// 	callback:    commandMapB,
		// },
	}

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		inputFields := strings.Fields(input)
		if len(inputFields) == 0 {
			continue
		}
		if obj, ok := commands[inputFields[0]]; ok {
			obj.callback()
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
