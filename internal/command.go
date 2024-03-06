package internal

import (
	"fmt"
	"os"
)

type command struct {
	Name        string
	Description string
	Callback    func()
}

func GetCommands() map[string]command {
	return map[string]command{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Get maps forward.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get maps backward.",
			Callback:    commandMapb,
		},
	}
}

func commandHelp() {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := GetCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Printf("\n")
}

func commandExit() {
	os.Exit(0)
}

func commandMap() {
	response, err := GetNextMaps()
	handleMapResponse(response, err)
}

func commandMapb() {
	response, err := GetPreviousMaps()
	handleMapResponse(response, err)
}

func handleMapResponse(response *MapResponse, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	maps := response.Maps

	for _, m := range maps {
		fmt.Println(m.Name)
	}
}
