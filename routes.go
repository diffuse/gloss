package gloss

import "github.com/go-chi/chi"

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/counter/{counterId}", IncrementCounterById)
	r.Get("/counter/{counterId}", GetCounterById)

	return r
}
