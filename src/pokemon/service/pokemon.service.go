package service

import (
	"github.com/leoomi/GoWebService/src/pokemon/data"
	"github.com/leoomi/GoWebService/src/pokemon/models"
)

type pokemonServiceImpl struct {
	data data.PokemonData
}

type PokemonService interface {
	Get(pokedexNumber int) (models.Pokemon, error)
}

func New() PokemonService {
	var data = data.New()

	return &pokemonServiceImpl{
		data: data,
	}
}

func (service *pokemonServiceImpl) Get(pokedexNumber int) (models.Pokemon, error) {
	return service.data.GetPokemon(pokedexNumber)
}
