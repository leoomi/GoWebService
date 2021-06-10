package pokemon

type PokemonServiceMock struct {
	GetFn     func(int) (Pokemon, error)
	GetCalled bool
}

type ServiceMockParams struct {
	GetReturn Pokemon
	GetError  error
}

func NewServiceMock(params ServiceMockParams) PokemonService {
	return &PokemonServiceMock{
		GetFn: func(_ int) (Pokemon, error) {
			return params.GetReturn, params.GetError
		},
	}
}

func (mock *PokemonServiceMock) Get(pokedexNumber int) (Pokemon, error) {
	mock.GetCalled = true
	return mock.GetFn(pokedexNumber)
}
