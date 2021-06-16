package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

type Api struct {
	router *httprouter.Router
}

type Route struct {
	Path    string
	Method  string
	Handler httprouter.Handle
}

type Controller interface {
	Routes() []Route
	Middlewares() []Middleware
}

func NewApi() *Api {
	return &Api{
		router: httprouter.New(),
	}
}

func (api *Api) RegisterController(controller Controller) {
	for _, route := range controller.Routes() {
		handler := route.Handler
		for _, middleware := range controller.Middlewares() {
			handler = middleware(handler)
		}

		switch route.Method {
		case http.MethodGet:
			api.router.GET(route.Path, handler)
		case http.MethodPost:
			api.router.POST(route.Path, handler)
		}
	}
}

func (api *Api) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(":8080", api.router))
}
