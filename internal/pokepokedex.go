package internal

import "fmt"

func ShowPokemonsInsidePokedex() {
	pokemons := pokedex.pokemons

	if len(pokemons) == 0 {
		fmt.Printf("No pokemons yet! Use the catch command to have some!\n\n")
	} else {
		fmt.Println("Your pokedex:")
		for _, pokemon := range pokemons {
			fmt.Printf("- %s\n", pokemon.Name)
		}
		fmt.Printf("\n")
	}
}
