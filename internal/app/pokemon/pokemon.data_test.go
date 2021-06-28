package pokemon

import (
	"fmt"
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

func TestGetPokemonQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM pokemon").
		WithArgs(1).
		WillReturnError(fmt.Errorf("error"))

	var data = &PokemonDataPG{
		db: db,
	}

	var _, error = data.GetPokemon(1)

	if error == nil {
		t.Error("Errror not returned")
	}
}

func TestGetPokemonScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"pokedexNumber", "name"}).
		AddRow("Error", "Error")
	mock.ExpectQuery("SELECT (.+) FROM pokemon").WithArgs(1).WillReturnRows(rows)

	var data = &PokemonDataPG{
		db: db,
	}

	var _, error = data.GetPokemon(1)

	if error == nil {
		t.Error("Errror not returned")
	}
}

func TestInsertPokemon(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO pokemon").
		WithArgs(1, "Name").
		WillReturnResult(sqlmock.NewResult(1, 1))

	var data = &PokemonDataPG{
		db: db,
	}

	err = data.PostPokemon(Pokemon{
		Name:          "Name",
		PokedexNumber: 1,
	})

	if err != nil {
		t.Error("Error returned")
	}
}

func TestInsertPokemonError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO pokemon").
		WithArgs(1, "Name").
		WillReturnError(fmt.Errorf("error"))

	var data = &PokemonDataPG{
		db: db,
	}

	err = data.PostPokemon(Pokemon{
		Name:          "Name",
		PokedexNumber: 1,
	})

	if err == nil {
		t.Error("Error not returned")
	}
}
