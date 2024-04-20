package main

import "github.com/rafaelnacle/gopherdex/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	initializeRepl(&cfg)
}
