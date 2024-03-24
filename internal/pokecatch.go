package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

type Pokedex struct {
	pokemons map[string]Pokemon
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy any    `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices            []any  `json:"game_indices"`
	Height                 int    `json:"height"`
	HeldItems              []any  `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []any  `json:"past_abilities"`
	PastTypes     []any  `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault any `json:"front_default"`
				FrontFemale  any `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      any `json:"back_default"`
					BackGray         any `json:"back_gray"`
					BackTransparent  any `json:"back_transparent"`
					FrontDefault     any `json:"front_default"`
					FrontGray        any `json:"front_gray"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      any `json:"back_default"`
					BackGray         any `json:"back_gray"`
					BackTransparent  any `json:"back_transparent"`
					FrontDefault     any `json:"front_default"`
					FrontGray        any `json:"front_gray"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           any `json:"back_default"`
					BackShiny             any `json:"back_shiny"`
					BackShinyTransparent  any `json:"back_shiny_transparent"`
					BackTransparent       any `json:"back_transparent"`
					FrontDefault          any `json:"front_default"`
					FrontShiny            any `json:"front_shiny"`
					FrontShinyTransparent any `json:"front_shiny_transparent"`
					FrontTransparent      any `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      any `json:"back_default"`
					BackShiny        any `json:"back_shiny"`
					FrontDefault     any `json:"front_default"`
					FrontShiny       any `json:"front_shiny"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      any `json:"back_default"`
					BackShiny        any `json:"back_shiny"`
					FrontDefault     any `json:"front_default"`
					FrontShiny       any `json:"front_shiny"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault any `json:"front_default"`
					FrontShiny   any `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  any `json:"back_default"`
					BackShiny    any `json:"back_shiny"`
					FrontDefault any `json:"front_default"`
					FrontShiny   any `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  any `json:"back_default"`
					BackShiny    any `json:"back_shiny"`
					FrontDefault any `json:"front_default"`
					FrontShiny   any `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      any `json:"back_default"`
					BackFemale       any `json:"back_female"`
					BackShiny        any `json:"back_shiny"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      any `json:"back_default"`
						BackFemale       any `json:"back_female"`
						BackShiny        any `json:"back_shiny"`
						BackShinyFemale  any `json:"back_shiny_female"`
						FrontDefault     any `json:"front_default"`
						FrontFemale      any `json:"front_female"`
						FrontShiny       any `json:"front_shiny"`
						FrontShinyFemale any `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault any `json:"front_default"`
					FrontFemale  any `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     any `json:"front_default"`
					FrontFemale      any `json:"front_female"`
					FrontShiny       any `json:"front_shiny"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type CatchChance struct {
	probability int
	description string
}

var pokedex = Pokedex{
	pokemons: make(map[string]Pokemon),
}

func CatchPokemon(name string) {
	pokemon, err := getRemotePokemon(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	chance := getChanceFrom(pokemon.BaseExperience)
	fmt.Println(chance.description)

	pokemonName := strings.ToUpper(name[:1]) + name[1:]

	is_catch := rand.Intn(chance.probability) == 1
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if is_catch {
		fmt.Printf("%s was caught!\n", pokemonName)
		pokedex.save(*pokemon)
	} else {
		fmt.Printf("%s escaped!\n\n", pokemonName)
	}
}

// getRemotePokemonInformation: Get REMOTE pokemon info from API.
func getRemotePokemon(pokemon string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	if res.StatusCode == 404 {
		err := errors.New("This pokemon doesn't exist!")
		return nil, err
	}

	if res.StatusCode != 200 {
		errString := fmt.Sprintf("Response failed.\nStatus code: %d\nBody: %s\n", res.StatusCode, body)
		err := errors.New(errString)
		return nil, err
	}

	response := Pokemon{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		error := errors.New("A problem happened when the program tries to decode the response.")
		return nil, error
	}

	return &response, nil
}

// CatchChance is based on the base_experience of a pokemon.
// Higher the base_experience is, less chance you have to catch it!
func getChanceFrom(baseExperience int) CatchChance {
	switch {
	case baseExperience <= 36:
		return CatchChance{1, "In your pocket!"}
	case baseExperience <= 50:
		return CatchChance{2, "Easy!"}
	case baseExperience <= 100:
		return CatchChance{4, "It's gonna be less easy"}
	case baseExperience <= 150:
		return CatchChance{8, "It's gonna be less easy"}
	case baseExperience <= 200:
		return CatchChance{16, "It's gonna be difficult"}
	case baseExperience <= 250:
		return CatchChance{32, "Very difficult!"}
	case baseExperience <= 290:
		return CatchChance{64, "You most likely won't have it!"}
	default:
		return CatchChance{128, "Good luck!"}
	}
}

// Save pokemon in the pokedex
func (pokedex *Pokedex) save(pokemon Pokemon) {
	mu.Lock()
	pokedex.pokemons[pokemon.Name] = pokemon
	mu.Unlock()
}
