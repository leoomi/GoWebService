package data

import (
	"database/sql"

	"github.com/leoomi/GoWebService/src/database"
	models "github.com/leoomi/GoWebService/src/pokemon/models"
	_ "github.com/lib/pq"
)

type pokemon = models.Pokemon
type PokemonData interface {
	GetPokemon(pokedexNumber int) (pokemon, error)
	PostPokemon(pokemon pokemon)
}

type PokemonDataPG struct {
	db *sql.DB
}

func New() PokemonData {
	return &PokemonDataPG{
		db: database.Pool,
	}
}

func (data *PokemonDataPG) GetPokemon(pokedexNumber int) (pokemon, error) {
	row := data.db.QueryRow("SELECT * FROM pokemon where pokedexNumber = $1", pokedexNumber)
	if row.Err() != nil {
		return pokemon{}, row.Err()
	}

	var pokemon pokemon
	err := row.Scan(&pokemon.PokedexNumber, &pokemon.Name)
	if err != nil {
		return models.Pokemon{}, err
	}

	return pokemon, nil
}

func (data *PokemonDataPG) PostPokemon(pokemon pokemon) {
	var err error
	sqlStatement := `
INSERT INTO users (pokedexNumber, name)
VALUES ($1, $2)`
	_, err = data.db.Exec(sqlStatement, pokemon.PokedexNumber, pokemon.Name)
	if err != nil {
		panic(err)
	}
}
