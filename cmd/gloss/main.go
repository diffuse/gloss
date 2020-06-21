package main

import (
	"github.com/diffuse/gloss"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	// close the db connections on exit
	defer gloss.Db.Close()

	// setup router
	r := chi.NewRouter()

	// set middleware
	r.Use(
		middleware.Logger,
		middleware.Recoverer)

	// mount routes on versioned path
	r.Mount("/v1", gloss.GetRoutes())

	// create server with some reasonable defaults
	s := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      5 * time.Second,
		MaxHeaderBytes:    64e3,
	}

	// serve
	log.Fatal(s.ListenAndServe())
}
