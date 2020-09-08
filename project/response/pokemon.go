package response

import (
	"github.com/google/uuid"

	"pahamify/project/model"
)

// Pokemon response
type Pokemon struct {
	ID     uuid.UUID `json:"id"`
	Number string    `json:"number"`
	Name   string    `json:"name"`
	Types  []string  `json:"types"`
}

// FromPokemonModel returns Pokemon response from model.Pokemon
func FromPokemonModel(pokemon model.Pokemon) (resp Pokemon) {
	resp.ID = pokemon.ID
	resp.Name = pokemon.Name
	resp.Number = pokemon.Number
	for _, pokemonType := range pokemon.Types {
		resp.Types = append(resp.Types, pokemonType.Name)
	}
	return
}

// FromPokemonsModel returns array of Pokemon response from model.Pokemons
func FromPokemonsModel(pokemons model.Pokemons) (resp []Pokemon) {
	for _, pokemon := range pokemons {
		pokemonResp := FromPokemonModel(pokemon)
		resp = append(resp, pokemonResp)
	}
	return
}
