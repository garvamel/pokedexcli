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

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	return location

}
