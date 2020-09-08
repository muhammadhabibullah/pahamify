package service

import (
	"pahamify/project/repository"
)

// Container contains services
type Container struct {
	PokemonService
}

// Init service Container
func Init(repo *repository.Container) *Container {
	return &Container{
		PokemonService: NewPokemonService(repo),
	}
}
