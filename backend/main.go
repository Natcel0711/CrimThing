package main

import (
	"crim/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", handler.GetProperty)

	http.ListenAndServe(":3000", r)
}
