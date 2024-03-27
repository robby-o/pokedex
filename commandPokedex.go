package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon yet")
	}

	fmt.Println("Your pokedex:")
	for _, v := range cfg.caughtPokemon {
		fmt.Println("  -", v.Name)
	}

	return nil
}
