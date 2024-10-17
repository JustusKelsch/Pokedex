package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config, args []string) error {
	locationResp, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.prevLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(config *config, args []string) error {

	if config.prevLocationsURL == nil {
		return errors.New("there are no previous locations")
	}
	
	locationResp, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResp.Next
	config.prevLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}