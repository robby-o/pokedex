package main

import (
	"time"

	"github.com/robby-o/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	startRepl(cfg)
}
