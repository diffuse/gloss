package chi

import (
	"github.com/diffuse/gloss"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"strconv"
)

type Router struct {
	*chi.Mux
	db gloss.Database
}

// NewRouter creates a router, associates a database
// with it, and mounts API routes
func NewRouter(db gloss.Database) *Router {
	r := &Router{}
	r.setupRoutes()
	r.SetDb(db)

	return r
}

// setupRoutes
func (rt *Router) setupRoutes() {
	// set routes
	routes := chi.NewRouter()
	routes.Post("/counter/{counterId}", rt.IncrementCounterById)
	routes.Get("/counter/{counterId}", rt.GetCounterById)

	// setup router
	rt.Mux = chi.NewRouter()

	// set middleware
	rt.Mux.Use(
		middleware.Logger,
		middleware.Recoverer)

	// mount routes on versioned path
	rt.Mux.Mount("/v1", routes)
}

// SetDb associates a gloss.Database with this router
func (rt *Router) SetDb(db gloss.Database) {
	rt.db = db
}

// Increment the value of the counter with ID counterId
func (rt *Router) IncrementCounterById(w http.ResponseWriter, r *http.Request) {
	// get the counter ID
	counterId, err := strconv.ParseUint(chi.URLParam(r, "counterId"), 10, 32)
	if err != nil {
		http.Error(w, "invalid counter ID", http.StatusBadRequest)
		return
	}

	// increment in db
	if err := rt.db.IncrementCounter(int(counterId)); err != nil {
		http.Error(w, "failed to increment counter value", http.StatusInternalServerError)
	}
}

// Get the value of the counter with ID counterId
func (rt *Router) GetCounterById(w http.ResponseWriter, r *http.Request) {
	// get the counter ID
	counterId, err := strconv.ParseInt(chi.URLParam(r, "counterId"), 10, 32)
	if err != nil {
		http.Error(w, "invalid counter ID", http.StatusBadRequest)
		return
	}

	// get value and return to requester
	val, err := rt.db.GetCounterVal(int(counterId))
	if err != nil {
		http.Error(w, "failed to get counter value", http.StatusNotFound)
		return
	}

	if _, err := w.Write([]byte(strconv.Itoa(val))); err != nil {
		panic(err)
	}
}
