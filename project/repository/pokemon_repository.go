package repository

import (
	"context"

	"gorm.io/gorm"

	"pahamify/project/model"
)

// PokemonRepository interface
type PokemonRepository interface {
	// CreatePokemon insert a new pokemon data
	CreatePokemon(ctx context.Context, pokemon *model.Pokemon) error
	// UpdatePokemon update an existing pokemon data
	UpdatePokemon(ctx context.Context, pokemon *model.Pokemon) error
	// GetPokemons retrieve pokemon data
	GetPokemons(ctx context.Context, limit int) (model.Pokemons, error)
	// DeletePokemon remove an existing pokemon data
	DeletePokemon(ctx context.Context, pokemon *model.Pokemon) error
}

type pokemonRepository struct {
	db *gorm.DB
}

// NewPokemonRepository returns PokemonRepository
func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{
		db: GetMySQL(),
	}
}

func (repo *pokemonRepository) CreatePokemon(ctx context.Context, pokemon *model.Pokemon) error {
	return repo.db.WithContext(ctx).
		Create(pokemon).Error
}

func (repo *pokemonRepository) UpdatePokemon(ctx context.Context, pokemon *model.Pokemon) error {
	return repo.db.WithContext(ctx).
		Where("id = ?", pokemon.ID).
		Updates(pokemon).Error
}

func (repo *pokemonRepository) GetPokemons(ctx context.Context, limit int) (
	pokemons model.Pokemons, err error) {
	err = repo.db.WithContext(ctx).
		Preload("Types").
		Limit(limit).
		Find(&pokemons).Error
	return
}

func (repo *pokemonRepository) DeletePokemon(ctx context.Context, pokemon *model.Pokemon) (err error) {
	tx := repo.db.WithContext(ctx).Begin()

	err = tx.Model(pokemon).
		Association("Types").
		Clear()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(pokemon).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
