package api

import "net/http"

type Middleware func(http.Handler) http.Handler

type Api struct {
	Mux      *http.ServeMux
	BasePath string
}

type Controller interface {
	Path() string
	Middlewares() []Middleware
	Handler() http.Handler
}

func NewApi() *Api {
	return &Api{
		Mux: http.NewServeMux(),
	}
}

func (api *Api) RegisterController(controller Controller) {
	handler := controller.Handler()
	for _, middleware := range controller.Middlewares() {
		handler = middleware(handler)
	}

	api.Mux.Handle(controller.Path(), handler)
}

func (api *Api) ListenAndServe(port string) {
	http.ListenAndServe(port, api.Mux)
}
