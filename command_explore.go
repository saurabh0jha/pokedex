package main

import "fmt"

func commandExplore(cfg *config, locationName string) error {
	resp, err := cfg.pokeapiClient.ExploreLocation(locationName)

	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
