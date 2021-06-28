package pokemon

type PokemonServiceMock struct {
	GetFn      func(int) (Pokemon, error)
	GetCalled  bool
	PostFn     func(_ Pokemon) error
	PostCalled bool
}

type ServiceMockParams struct {
	GetReturn Pokemon
	GetError  error
	PostError error
}

func NewServiceMock(params ServiceMockParams) PokemonService {
	return &PokemonServiceMock{
		GetFn: func(_ int) (Pokemon, error) {
			return params.GetReturn, params.GetError
		},
		PostFn: func(_ Pokemon) error {
			return params.PostError
		},
	}
}

func (mock *PokemonServiceMock) Get(pokedexNumber int) (Pokemon, error) {
	mock.GetCalled = true
	return mock.GetFn(pokedexNumber)
}

func (mock *PokemonServiceMock) Post(pokemon Pokemon) error {
	mock.PostCalled = true
	return mock.PostFn(pokemon)
}
