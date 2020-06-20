package gloss

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// Increment the value of the counter with ID counterId
func IncrementCounterById(w http.ResponseWriter, r *http.Request) {
	// get the counter ID
	counterId, err := strconv.ParseUint(chi.URLParam(r, "counterId"), 10, 32)
	if err != nil {
		http.Error(w, "invalid counter ID", http.StatusBadRequest)
		return
	}

	// increment in db
	if err := IncrementCounter(int(counterId)); err != nil {
		http.Error(w, "failed to increment counter value", http.StatusInternalServerError)
	}
}

// Get the value of the counter with ID counterId
func GetCounterById(w http.ResponseWriter, r *http.Request) {
	// get the counter ID
	counterId, err := strconv.ParseUint(chi.URLParam(r, "counterId"), 10, 32)
	if err != nil {
		http.Error(w, "invalid counter ID", http.StatusBadRequest)
		return
	}

	// get value and return to requester
	val, err := GetCounterVal(int(counterId))
	if err != nil {
		http.Error(w, "failed to get counter value", http.StatusInternalServerError)
	}

	if _, err := w.Write([]byte(strconv.Itoa(val))); err != nil {
		panic(err)
	}
}
