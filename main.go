package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/garvamel/pokedexcli/internal/pokeapi"
)

type configuration struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration) error
}

var commands map[string]cliCommand

func main() {

	conf := configuration{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}

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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapback",
			description: "Lists previous location areas",
			callback:    commandMapBack,
		},
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
			obj.callback(&conf)
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func commandExit(conf *configuration) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *configuration) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *configuration) error {

	loc := pokeapi.GetLocation(conf.next)

	conf.next = loc.Next
	s, _ := loc.Previous.(string)
	conf.previous = s
	return nil
}

func commandMapBack(conf *configuration) error {

	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	loc := pokeapi.GetLocation(conf.previous)

	conf.next = loc.Next
	s, _ := loc.Previous.(string)
	conf.previous = s
	return nil
}
