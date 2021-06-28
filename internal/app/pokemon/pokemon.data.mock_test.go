package pokemon

type PokemonDataMock struct {
	GetFn     func(int) (Pokemon, error)
	GetCalled bool

	PostFn     func(Pokemon) error
	PostCalled bool
}

type DataMockParams struct {
	GetReturn Pokemon
	GetError  error
	PostError error
}

func NewDataMock(params DataMockParams) pokemonData {
	return &PokemonDataMock{
		GetFn: func(_ int) (Pokemon, error) {
			return params.GetReturn, params.GetError
		},
		PostFn: func(_ Pokemon) error {
			return params.PostError
		},
	}
}

func (mock *PokemonDataMock) GetPokemon(pokedexNumber int) (Pokemon, error) {
	mock.GetCalled = true
	return mock.GetFn(pokedexNumber)
}

func (mock *PokemonDataMock) PostPokemon(pokemon Pokemon) error {
	mock.PostCalled = true
	return mock.PostFn(pokemon)
}
