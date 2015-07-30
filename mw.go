// Package mw provides a Decorate function to cleanly decorate
// http.Handler interfaces with multiple middlewares.
//
/*
Go http middleware uses the following form

	func Middleware(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// middleware logic

			h.ServeHTTP(w, r)
		})
	}

You can use it with any existing middleware.

	import (
		// ...
		"github.com/collinglass/mw"
		"github.com/justinas/nosurf"
	)

	// decorate router
	server := mw.Decorate(
		http.NewServeMux(),
		// add middleware from existing packages
		nosurf.NewPure,
	)

Or you can build your own custom middleware.

	import (
		// ...
		"github.com/collinglass/mw"
	)

	// Create middleware from scratch
	func JSONMiddleware(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set Response Content-Type to application/json
			w.Header().Set("Content-Type", "application/json")

			h.ServeHTTP(w, r)
		})
	}

You can use it with gorilla mux.

	// new router
	r := mux.NewRouter()

	r.HandleFunc("/api/data", DataHandler).Methods("GET")

	// decorate router
	server := mw.Decorate(
		r,
		nosurf.NewPure,
		JSONMiddleware,
	)

	http.Handle("/api/", server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

Or with the standard library http.ServeMux.

	// new router
	r := http.NewServeMux()

	r.HandleFunc("/api/data", DataHandler)

	// decorate router
	server := mw.Decorate(
		r,
		nosurf.NewPure,
		JSONMiddleware,
	)

	http.Handle("/api/", server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

*/
package mw

import (
	"net/http"
)

// Middleware takes an http.Handler interface and decorates it.
type Middleware func(http.Handler) http.Handler

// Decorate ranges over a variadic number of middleware and
// decorates the http.Handler with them.
func Decorate(h http.Handler, ds ...Middleware) http.Handler {
	decorated := h
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
