package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	roll := rand.Intn(pokemon.BaseExperience)
	threshold := 20
	if roll > threshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught \n", pokemon.Name)
		cfg.pokeDex[pokemon.Name] = pokemon
	}
	return nil

}
