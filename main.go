package main

import (
	"time"

	"github.com/saurabh0jha/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 20*time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
