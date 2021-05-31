package data

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetPokemon(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"pokedexNumber", "name"}).
		AddRow(1, "Name 1").
		AddRow(2, "Name 2")
	mock.ExpectQuery("SELECT (.+) FROM pokemon").WithArgs(1).WillReturnRows(rows)

	var data = &PokemonDataPG{
		db: db,
	}

	var result, _ = data.GetPokemon(1)

	if result.Name != "Name 1" {
		t.Error("Wrong name")
	}
}
