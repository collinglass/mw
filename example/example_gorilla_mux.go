package main

import (
	"github.com/collinglass/mw"
	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"data":"json"}`))
}

// Create middleware from scratch
func JSONMiddleware() mw.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set JSON Response Header
			w.Header().Set("Content-Type", "application/json")

			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	// new router
	r := mux.NewRouter()

	r.HandleFunc("/api/data", DataHandler).Methods("GET")

	// decorate router
	server := mw.Decorate(
		r,
		// add middleware from existing packages
		nosurf.NewPure,
		// or your own
		JSONMiddleware(),
	)

	http.Handle("/api/", server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
