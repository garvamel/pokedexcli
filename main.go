package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	callback    func(*configuration, string) error
}

var commands map[string]cliCommand

var pokedex map[string]pokeapi.Pokemon

func main() {

	pokedex = map[string]pokeapi.Pokemon{}

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
		"explore": {
			name:        "explore",
			description: "Lists pokemon found in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Check Pokedex for entry",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught Pokemon",
			callback:    commandPokedex,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		fields := cleanInput(scanner.Text())

		if l := len(fields); l == 0 {
			continue
		} else if l == 1 {
			if obj, ok := commands[fields[0]]; ok {
				obj.callback(&conf, "")
			} else {
				fmt.Println("Unknown command")
			}
		} else {
			if obj, ok := commands[fields[0]]; ok {
				obj.callback(&conf, fields[1])
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	inputFields := strings.Fields(input)
	return inputFields
}

func commandExit(conf *configuration, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *configuration, arg string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(conf *configuration, arg string) error {

	loc := pokeapi.GetLocation(conf.next)

	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}

	conf.next = loc.Next
	s, _ := loc.Previous.(string)
	conf.previous = s
	return nil
}

func commandMapBack(conf *configuration, arg string) error {

	if conf.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	loc := pokeapi.GetLocation(conf.previous)

	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}

	conf.next = loc.Next
	s, _ := loc.Previous.(string)
	conf.previous = s
	return nil
}

func commandExplore(conf *configuration, arg string) error {

	if arg == "" {
		fmt.Println("No area to explore")
		return nil
	}
	area := pokeapi.GetAreaPokemon(arg)

	fmt.Println("Exploring", arg+"...")
	fmt.Println("Found Pokemon:")

	for _, result := range area.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *configuration, arg string) error {

	if arg == "" {
		fmt.Println("No area to explore")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	pokemon := pokeapi.GetPokemonByName(arg)

	catchChance := rand.Intn(609) // 608 is the highest base experience

	if catchChance >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func commandInspect(conf *configuration, arg string) error {

	if arg == "" {
		fmt.Println("no pokemon to inspect")
		return nil
	}

	pkmn, ok := pokedex[arg]

	if ok {
		fmt.Printf("Name: %s\n", pkmn.Name)
		fmt.Printf("Height: %s\n", pkmn.Name)
		fmt.Printf("Weight: %s\n", pkmn.Name)
		fmt.Println("Stats:")
		for _, stat := range pkmn.Stats {

			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Type:")
		for _, pktype := range pkmn.Types {

			fmt.Printf("  -%s\n", pktype.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil

}

func commandPokedex(conf *configuration, arg string) error {

	fmt.Println("Your Pokemon:")

	if pokedex != nil {
		for pkmn := range pokedex {
			fmt.Printf("  - %v\n", pkmn)
		}

	}

	return nil
}
