package internal

import "fmt"

func InspectPokemon(name string) {
	if pokemon, ok := pokedex.pokemons[name]; ok {
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("- %s", t.Type.Name)
		}

		fmt.Printf("\n\n")
	} else {
		fmt.Printf("you have not caught that pokemon\n\n")
	}
}
