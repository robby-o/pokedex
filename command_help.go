package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	for _, cmd := range getCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
