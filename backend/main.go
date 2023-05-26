package main

import (
	"net/http"
	"crim/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Property struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Catastro string `json:"catastro"`
	DireccionFisica string `json:"direccion_fisica"`
	DireccionPostal string `json:"direccion_postal"`
}

func main(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", handler.GetProperty)

	http.ListenAndServe(":3000", r)
}


