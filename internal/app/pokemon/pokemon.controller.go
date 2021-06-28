package pokemon

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

func (controller *pokemonController) Routes() []api.Route {
	return []api.Route{{
		Path:    "/pokemon/:id",
		Method:  http.MethodGet,
		Handler: controller.get,
	}, {
		Path:    "/pokemon",
		Method:  http.MethodPost,
		Handler: controller.post,
	}}
}

func (controller *pokemonController) get(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	idString := ps.ByName("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	pokemon, err := controller.service.Get(id)

	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	pokemonJSON, _ := json.Marshal(pokemon)
	res.Write(pokemonJSON)
}

func (controller *pokemonController) post(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	controller.service.Post(pokemon)
	res.Write([]byte("banana"))
}
