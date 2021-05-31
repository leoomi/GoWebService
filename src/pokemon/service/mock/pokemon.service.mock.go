package mock

import (
	"github.com/leoomi/GoWebService/src/pokemon/models"
	"github.com/leoomi/GoWebService/src/pokemon/service"
)

type PokemonServiceMock struct {
	GetFn     func(int) (models.Pokemon, error)
	GetCalled bool
}

type ServiceMockParams struct {
	GetReturn models.Pokemon
	GetError  error
}

func NewServiceMock(params ServiceMockParams) service.PokemonService {
	return &PokemonServiceMock{
		GetFn: func(_ int) (models.Pokemon, error) {
			return params.GetReturn, params.GetError
		},
	}
}

func (mock *PokemonServiceMock) Get(pokedexNumber int) (models.Pokemon, error) {
	mock.GetCalled = true
	return mock.GetFn(pokedexNumber)
}
