package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/leoomi/GoWebService/src/api"
)

const path = "/pokemon"

type pokemonService struct{}

func New() api.Service {
	return pokemonService{}
}

func (pokemonService) Path() string {
	return path
}

func (pokemonService) Middlewares() []api.Middleware {
	return make([]api.Middleware, 0)
}

func (pokemonService) Handler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			post(res, req)
		case http.MethodGet:
			get(res, req)
		case http.MethodPut:
			get(res, req)
		}
	})
}

func get(res http.ResponseWriter, req *http.Request) {
	pokemon := GetPokemon(1)

	pokemonJSON, _ := json.Marshal(pokemon)
	res.Write([]byte(pokemonJSON))
}

func post(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("banana"))
}
