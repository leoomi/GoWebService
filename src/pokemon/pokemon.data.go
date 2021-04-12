package pokemon

import (
	db "github.com/leoomi/GoWebService/src/database"
	_ "github.com/lib/pq"
)

func GetPokemon(pokedexNumber int) Pokemon {
	row := db.Pool.QueryRow("SELECT pokedexnumber, name FROM pokemon")
	if row.Err() != nil {
		panic(row.Err().Error())
	}

	var pokemon Pokemon
	err := row.Scan(&pokemon.PokedexNumber, &pokemon.Name)
	if err != nil {
		panic(err.Error())
	}

	return pokemon
}
