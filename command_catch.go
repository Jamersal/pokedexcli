package main

import "fmt"

func commandCatch(cfg *config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	return nil

}
