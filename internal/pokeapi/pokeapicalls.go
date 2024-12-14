package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/garvamel/pokedexcli/internal/pokecache"
)

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type exploreArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func GetLocation(url string) location {

	var body []byte
	var err error
	var res *http.Response

	c := pokecache.NewCache(5)

	val, ok := c.Get(url)

	if ok {
		body = val
	} else {

		res, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)

		}
	}

	location := location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	return location

}

func GetAreaPokemon(area string) exploreArea {

	url := "https://pokeapi.co/api/v2/location-area/" + area

	var body []byte
	var err error
	var res *http.Response

	c := pokecache.NewCache(5)

	val, ok := c.Get(url)

	if ok {
		body = val
	} else {

		res, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)

		}
	}

	expArea := exploreArea{}
	err = json.Unmarshal(body, &expArea)
	if err != nil {
		fmt.Println(err)
	}

	return expArea
}

func GetPokemonByName(name string) Pokemon {

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	var body []byte
	var err error
	var res *http.Response

	res, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)

	}

	poke := Pokemon{}
	err = json.Unmarshal(body, &poke)
	if err != nil {
		fmt.Println(err)
	}

	return poke
}
