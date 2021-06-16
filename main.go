package main

import (
	"net/http"

	"github.com/leoomi/GoWebService/internal/app/pokemon"
	"github.com/leoomi/GoWebService/internal/pkg/database"
	"github.com/leoomi/GoWebService/pkg/api"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//before
		handler.ServeHTTP(res, req)
		//after
	})
}

func main() {
	database.Init()
	defer database.Close()
	pokemonApi := api.NewApi()
	pokemonApi.RegisterController(pokemon.New())
	pokemonApi.ListenAndServe(":8080")
}
