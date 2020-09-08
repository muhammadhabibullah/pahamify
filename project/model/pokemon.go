package model

// Pokemon model
type Pokemon struct {
	GormModel
	Number string `gorm:"number" json:"number"`
	Name   string `gorm:"name" json:"name"`
	Types  Types  `gorm:"many2many:pokemon_types;" json:"types"`
}

// Pokemons alias of array of Pokemon
type Pokemons []Pokemon

// Type model
type Type struct {
	Name string `gorm:"type:varchar(256);primary_key;name"`
}

// Types alias of array of Type
type Types []Type
