package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Pokemon struct {
	name           string
	baseExperience int
	id             int
	height         int
	order          int
	weight         int
}

func commandCatch(cfg *config, pokemonName string) error {

	pokemonName = strings.ToLower(pokemonName)

	if pokemonName == "" {
		err := fmt.Errorf("Please enter the name of Pokemon to catch")
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	resp, err := cfg.pokeapiClient.PokemonDetails(pokemonName)

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	chance := float64(rand.Intn(resp.BaseExperience)) / float64(resp.BaseExperience)
	fmt.Printf("chance %v \n", chance)

	if float64(chance) >= 0.5 {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = Pokemon{
			baseExperience: resp.BaseExperience,
			height:         resp.Height,
			id:             resp.ID,
			name:           pokemonName,
			order:          resp.Order,
			weight:         resp.Weight,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
