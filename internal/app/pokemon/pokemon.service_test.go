package pokemon

import (
	"errors"
	"testing"
)

func TestServiceGetPokemon(t *testing.T) {
	var mockParams = DataMockParams{
		GetReturn: Pokemon{
			PokedexNumber: 0,
			Name:          "Name",
		},
		GetError: nil,
	}

	var service = pokemonServiceImpl{
		data: NewDataMock(mockParams),
	}

	var result, _ = service.Get(1)

	if result.PokedexNumber != 0 {
		t.Error("Incorrect podexNumber")
	}

	if result.Name != "Name" {
		t.Error("Incorrect name")
	}
}

func TestGetPokemon2(t *testing.T) {
	var mockParams = DataMockParams{
		GetReturn: Pokemon{},
		GetError:  errors.New("error"),
	}

	var service = pokemonServiceImpl{
		data: NewDataMock(mockParams),
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
