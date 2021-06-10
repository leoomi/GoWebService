package pokemon

import (
	"database/sql"

	"github.com/leoomi/GoWebService/internal/pkg/database"
	_ "github.com/lib/pq"
)

type pokemonData interface {
	GetPokemon(pokedexNumber int) (Pokemon, error)
	PostPokemon(pokemon Pokemon)
}

type PokemonDataPG struct {
	db *sql.DB
}

func newData() pokemonData {
	return &PokemonDataPG{
		db: database.Pool,
	}
}

func (data *PokemonDataPG) GetPokemon(pokedexNumber int) (Pokemon, error) {
	row := data.db.QueryRow("SELECT * FROM pokemon where pokedexNumber = $1", pokedexNumber)
	if row.Err() != nil {
		return Pokemon{}, row.Err()
	}

	var pokemon Pokemon
	err := row.Scan(&pokemon.PokedexNumber, &pokemon.Name)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (data *PokemonDataPG) PostPokemon(pokemon Pokemon) {
	var err error
	sqlStatement := `
INSERT INTO users (pokedexNumber, name)
VALUES ($1, $2)`
	_, err = data.db.Exec(sqlStatement, pokemon.PokedexNumber, pokemon.Name)
	if err != nil {
		panic(err)
	}
}
