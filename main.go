package main

import (
	"log"
	"net/http"

	"github.com/ValenCedano/MinibackendHeroes/handlers"
	"github.com/ValenCedano/MinibackendHeroes/repositorio"
)

func main() {
	bd := repositorio.NewBaseDatos()
	handler := handlers.NewHandlerSuperheroes(bd)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/superhero", handler.GetSuperhero())

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
