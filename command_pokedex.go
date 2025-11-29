package main

import "fmt"

func commandPokedex(cfg *config, locationName string) error {

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s \n", pokemon.name)
	}

	return nil
}
