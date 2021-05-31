package service

import (
	"errors"
	"testing"

	"github.com/leoomi/GoWebService/src/pokemon/data/mock"
	"github.com/leoomi/GoWebService/src/pokemon/models"
)

func TestGetPokemon(t *testing.T) {
	var mockParams = mock.DataMockParams{
		GetReturn: models.Pokemon{
			PokedexNumber: 0,
			Name:          "Name",
		},
		GetError: nil,
	}

	var service = pokemonServiceImpl{
		data: mock.NewDataMock(mockParams),
	}

	var result, _ = service.Get(1)

	if result.PokedexNumber != 1 {
		t.Error("Incorrect podexNumber")
	}

	if result.Name != "Name" {
		t.Error("Incorrect name")
	}
}

func TestGetPokemon2(t *testing.T) {
	var mockParams = mock.DataMockParams{
		GetReturn: models.Pokemon{},
		GetError:  errors.New("error"),
	}

	var service = pokemonServiceImpl{
		data: mock.NewDataMock(mockParams),
	}

	var result, error = service.Get(1)

	if error == nil {
		t.Error("Error not returned")
	}

	if result.PokedexNumber != 0 {
		t.Error("Incorrect podexNumber")
	}

	if result.Name != "" {
		t.Error("Incorrect name")
	}
}
