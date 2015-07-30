# mw [![Coverage Status](https://coveralls.io/repos/collinglass/mw/badge.png?branch=master)](https://coveralls.io/r/collinglass/mw?branch=master) [![GoDoc](https://godoc.org/github.com/collinglass/mw?status.png)](https://godoc.org/github.com/collinglass/mw)

mw is a middleware decorator for go servers.

## Getting started

To start using mw, install Go and run `go get`:

```sh
$ go get github.com/collinglass/mw
```

### Using it with existing middleware

``` go
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
```

### Building your own middleware

``` go
import (
	// ...
	"github.com/collinglass/mw"
)

// Create middleware from scratch
func JSONMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Response Header "Content-Type" to JSON
		w.Header().Set("Content-Type", "application/json")

		h.ServeHTTP(w, r)
	})
}
```

### Using it with gorilla mux

``` go
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
```

### Using it with http.ServeMux

``` go
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
```

