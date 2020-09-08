package service

import (
	"context"
	"time"

	"pahamify/project/config"
	"pahamify/project/model"
	"pahamify/project/repository"
)

// PokemonService interface
type PokemonService interface {
	// CreatePokemon insert a new pokemon data
	CreatePokemon(context.Context, *model.Pokemon) error
	// UpdatePokemon update an existing pokemon data
	UpdatePokemon(context.Context, *model.Pokemon) error
	// GetPokemons retrieve pokemon data
	GetPokemons(context.Context, int) (model.Pokemons, error)
	// DeletePokemon remove an existing pokemon data
	DeletePokemon(context.Context, *model.Pokemon) error
}

type pokemonService struct {
	pokemonRepository repository.PokemonRepository
	processTimeout    int
}

// NewPokemonService returns PokemonService
func NewPokemonService(repo *repository.Container) PokemonService {
	return &pokemonService{
		pokemonRepository: repo.PokemonRepository,
		processTimeout:    config.GetConfig().Server.RequestTimeout,
	}
}

func (svc *pokemonService) setServiceTimeout(c context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c, time.Duration(svc.processTimeout)*time.Second)
	return ctx, cancel
}

func (svc *pokemonService) CreatePokemon(c context.Context, pokemon *model.Pokemon) error {
	ctx, cancel := svc.setServiceTimeout(c)
	defer cancel()

	return svc.pokemonRepository.CreatePokemon(ctx, pokemon)
}

func (svc *pokemonService) UpdatePokemon(c context.Context, pokemon *model.Pokemon) error {
	ctx, cancel := svc.setServiceTimeout(c)
	defer cancel()

	return svc.pokemonRepository.UpdatePokemon(ctx, pokemon)
}

func (svc *pokemonService) GetPokemons(c context.Context, limit int) (model.Pokemons, error) {
	ctx, cancel := svc.setServiceTimeout(c)
	defer cancel()

	return svc.pokemonRepository.GetPokemons(ctx, limit)
}

func (svc *pokemonService) DeletePokemon(c context.Context, pokemon *model.Pokemon) error {
	ctx, cancel := svc.setServiceTimeout(c)
	defer cancel()

	return svc.pokemonRepository.DeletePokemon(ctx, pokemon)
}
