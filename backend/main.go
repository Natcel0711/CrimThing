package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{where}", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println(chi.URLParam(r, "where"))
		w.Write([]byte("{ \"PUTA\": 1 }"))
	})

	http.ListenAndServe(":3000", r)
}

