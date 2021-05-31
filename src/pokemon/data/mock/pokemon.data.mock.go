package mock

import (
	"github.com/leoomi/GoWebService/src/pokemon/data"
	"github.com/leoomi/GoWebService/src/pokemon/models"
)

type PokemonDataMock struct {
	GetFn     func(int) (models.Pokemon, error)
	GetCalled bool

	PostFn     func(models.Pokemon)
	PostCalled bool
}

type DataMockParams struct {
	GetReturn models.Pokemon
	GetError  error
}

func NewDataMock(params DataMockParams) data.PokemonData {
	return &PokemonDataMock{
		GetFn: func(_ int) (models.Pokemon, error) {
			return params.GetReturn, params.GetError
		},
		PostFn: func(_ models.Pokemon) {},
	}
}

func (mock *PokemonDataMock) GetPokemon(pokedexNumber int) (models.Pokemon, error) {
	mock.GetCalled = true
	return mock.GetFn(pokedexNumber)
}

func (mock *PokemonDataMock) PostPokemon(pokemon models.Pokemon) {
	mock.PostCalled = true
	mock.PostFn(pokemon)
}
