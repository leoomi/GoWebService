package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/leoomi/GoWebService/pkg/api"
)

type pokemonController struct {
	service PokemonService
}

func New() api.Controller {
	var service = newService()

	return &pokemonController{
		service: service,
	}
}

func (*pokemonController) Path() string {
	return "/pokemon"
}

func (*pokemonController) Middlewares() []api.Middleware {
	return make([]api.Middleware, 0)
}

func (controller *pokemonController) Handler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			controller.post(res, req)
		case http.MethodGet:
			controller.get(res, req)
		case http.MethodPut:
			controller.get(res, req)
		}
	})
}

func (controller *pokemonController) get(res http.ResponseWriter, _ *http.Request) {
	pokemon, err := controller.service.Get(1)

	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	pokemonJSON, _ := json.Marshal(pokemon)
	res.Write(pokemonJSON)
}

func (*pokemonController) post(res http.ResponseWriter, _ *http.Request) {
	res.Write([]byte("banana"))
}
