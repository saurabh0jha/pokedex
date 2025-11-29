package main

import (
	"fmt"
)

func commandMapb(config *config, param string) error {
	if config.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)

	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.prevLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
