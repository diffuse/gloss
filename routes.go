package gloss

import "github.com/go-chi/chi"

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()

	return r
}
