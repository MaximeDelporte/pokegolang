package internal

import (
	"fmt"
	"os"
)

type command struct {
	Name string
	Description string
	Callback func() error
}

func GetCommands() map[string]command {
	return map[string]command{
		"help": {
			Name: 		 "help",
			Description: "Displays a help message",
			Callback: 	 commandHelp,
		},
		"exit": {
			Name: 		 "exit",
			Description: "Exit the Pokedex",
			Callback: 	 commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := GetCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Println("")

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

//https://pokeapi.co/api/v2/location/?offset=0&limit=20
