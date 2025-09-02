package main

import (
	"errors"
	"fmt"
	"log"
)

// callbackMap lists the next page of location areas using cfg.nextLocURL.
// On first invocation (nil nextLocURL) it fetches the first page.
func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocURL)
	if err != nil {
		log.Println("error fetching location areas:", err)
		return err
	}
	// Update pagination pointers.
	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("  -- %s\n", area.Name)
	}
	if cfg.nextLocURL == nil {
		fmt.Println("(end of list)")
	}
	return nil
}

// callbackMapb goes backwards (previous page). If no previous page it returns an error.
func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocURL)
	if err != nil {
		log.Println("error fetching previous location areas:", err)
		return err
	}
	// After moving backwards, set pagination pointers based on the response.
	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("  -- %s\n", area.Name)
	}
	return nil
}
