package repository

// Container contains repositories
type Container struct {
	PokemonRepository
}

// Init repository Container
func Init() *Container {
	return &Container{
		PokemonRepository: NewPokemonRepository(),
	}
}
