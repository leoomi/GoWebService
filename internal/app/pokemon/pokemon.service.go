package pokemon

type pokemonServiceImpl struct {
	data pokemonData
}

type PokemonService interface {
	Get(pokedexNumber int) (Pokemon, error)
	Post(pokemon Pokemon) error
}

func newService() PokemonService {
	var data = newData()

	return &pokemonServiceImpl{
		data: data,
	}
}

func (service *pokemonServiceImpl) Get(pokedexNumber int) (Pokemon, error) {
	return service.data.GetPokemon(pokedexNumber)
}

func (service *pokemonServiceImpl) Post(pokemon Pokemon) error {
	return service.data.PostPokemon(pokemon)
}
