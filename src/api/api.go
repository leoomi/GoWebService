package api

import "net/http"

type Middleware func(http.Handler) http.Handler

type Api struct {
	Mux      *http.ServeMux
	BasePath string
}

type Service interface {
	Path() string
	Middlewares() []Middleware
	Handler() http.Handler
}

func NewApi() Api {
	return Api{
		Mux: http.NewServeMux(),
	}
}

func (api Api) RegisterService(service Service) {
	handler := service.Handler()
	for _, middleware := range service.Middlewares() {
		handler = middleware(handler)
	}

	api.Mux.Handle(service.Path(), handler)
}

func (api Api) ListenAndServe(port string) {
	http.ListenAndServe(port, api.Mux)
}
