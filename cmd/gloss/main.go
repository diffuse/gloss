package main

import (
	"github.com/diffuse/gloss/chi"
	"github.com/diffuse/gloss/pgsql"
	"log"
	"net/http"
	"time"
)

func main() {
	// create a thread-safe database instance for use with the router
	log.Println("connecting to database")
	db := pgsql.NewDatabase()
	defer db.Close()

	// create a router and associate the database with it
	log.Println("setting up routes")
	r := chi.NewRouter(db)

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
	log.Println("serving")
	log.Fatal(s.ListenAndServe())
}
