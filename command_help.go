package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("X----Gopherdex Help Menu----X")
	fmt.Println("Available commands: ")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")

	return nil
}
