package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ValenCedano/MinibackendHeroes/repositorio"
)

type HandlerSuperheroes struct {
	BD *repositorio.BaseDatos
}

func NewHandlerSuperheroes(bd *repositorio.BaseDatos) *HandlerSuperheroes {
	return &HandlerSuperheroes{
		BD: bd,
	}
}

func (hs *HandlerSuperheroes) GetSuperhero() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		heroName := r.URL.Query().Get("hero")
		if heroName == "" {
			http.Error(w, "Se requiere ingresar el Nombre del Superheroe ", http.StatusBadRequest)
			return
		}

		hero, found := hs.BD.GetSuperheroByName(heroName)
		if !found {
			http.Error(w, "Superheroe No encontrado", http.StatusNotFound)
			return
		}

		payload, err := json.Marshal(hero)
		if err != nil {
			http.Error(w, "Fallo la codificacion a JSON del SuperHeroe", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	})
}
