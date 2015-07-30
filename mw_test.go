package mw_test

import (
	"github.com/collinglass/mw"
	"net/http"
	"testing"
)

// Create middleware from scratch
func CreatedMiddleware() mw.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set JSON Response Header
			w.WriteHeader(http.StatusCreated)

			h.ServeHTTP(w, r)
		})
	}
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

func TestDecorate(t *testing.T) {
	// new router
	r := http.NewServeMux()
	expectedContentType := "application/json"
	expectedStatus := http.StatusCreated

	r.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data":"json"}`))
	})

	// decorate router
	server := mw.Decorate(
		r,
		CreatedMiddleware(),
		JSONMiddleware(),
	)

	http.Handle("/api/", server)
	go http.ListenAndServe(":8080", nil)

	c := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/data", nil)
	if err != nil {
		t.Errorf("error creating request: %s\n", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		t.Errorf("error making request: %s\n", err)
	}

	// checks if JSONMiddlware was called
	actualContentType := resp.Header.Get("Content-Type")
	if expectedContentType != actualContentType {
		t.Errorf("expected %s and got %s", expectedContentType, actualContentType)
	}

	// checks if CreatedMiddleware was called
	actualStatus := resp.StatusCode
	if expectedStatus != actualStatus {
		t.Errorf("expected %s and got %s", expectedContentType, actualContentType)
	}
}
