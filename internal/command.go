package internal

import (
	"fmt"
	"os"
)

type command struct {
	Name        string
	Description string
	Callback    func(parameter any)
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
		"explore": {
			Name:        "explore [city_name]",
			Description: "Usage: explore city_name. Where a city_name is a value from map result. Returns a list of pokemon that can be found in this area.",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch [pokemon]",
			Description: "Usage: catch [pokemon] where pokemon is the name of the pokemon found in the list of the explore command",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect [pokemon]",
			Description: "Usage: inspect [pokemon] where pokemon is the name of the caught pokemon. You need to capture it before inspect it!",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Usage: pokedex. Use it to see all your caught pokemons!",
			Callback:    commandPokedex,
		},
	}
}

func commandHelp(parameter any) {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := GetCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Printf("\n")
}

func commandExit(parameter any) {
	os.Exit(0)
}

func commandMap(parameter any) {
	response, err := GetNextMaps()
	handleMapResponse(response, err)
}

func commandMapb(parameter any) {
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

func commandExplore(anyCity any) {
	city := anyCity.(string)
	response, err := GetPokemonsFrom(city)

	if err != nil {
		fmt.Println(err)
		return
	}

	pokemons := response.Pokemons
	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Pokemon.Name)
	}
}

func commandCatch(anyPokemon any) {
	pokemon := anyPokemon.(string)
	CatchPokemon(pokemon)
}

func commandInspect(anyPokemon any) {
	pokemon := anyPokemon.(string)
	InspectPokemon(pokemon)
}

func commandPokedex(parameter any) {
	ShowPokemonsInsidePokedex()
}
