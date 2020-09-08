package request

import (
	"github.com/google/uuid"

	"pahamify/project/model"
)

// Pokemon request
type Pokemon struct {
	ID     string   `json:"id"`
	Number string   `json:"number"`
	Name   string   `json:"name"`
	Types  []string `json:"types"`
}

// ToPokemonModel generate model.Pokemon from Pokemon request
func (req Pokemon) ToPokemonModel() model.Pokemon {
	reqUUID, _ := uuid.Parse(req.ID)
	pokemon := model.Pokemon{
		GormModel: model.GormModel{ID: reqUUID},
		Number:    req.Number,
		Name:      req.Name,
	}

	for _, t := range req.Types {
		pokemon.Types = append(pokemon.Types, model.Type{Name: t})
	}

	return pokemon
}
