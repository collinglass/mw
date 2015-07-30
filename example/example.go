package main

import (
	"github.com/collinglass/mw"
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
			return
		})
	}
}

func main() {
	// new router
	r := http.NewServeMux()

	r.HandleFunc("/api/data", DataHandler)

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
