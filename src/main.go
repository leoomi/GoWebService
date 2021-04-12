package main

import (
	"net/http"

	"github.com/leoomi/GoWebService/src/api"
	"github.com/leoomi/GoWebService/src/database"
	"github.com/leoomi/GoWebService/src/pokemon"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//before
		handler.ServeHTTP(res, req)
		//after
	})
}

func main() {
	pokemonApi := api.NewApi()
	pokemonApi.RegisterService(pokemon.New())
	pokemonApi.ListenAndServe(":8080")
	//log.Fatal(http.ListenAndServe(":8080", pokemonApi.Mux))

	database.Close()
}
