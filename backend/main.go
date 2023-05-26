package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{where}", func (w http.ResponseWriter, r *http.Request) {
		where := chi.URLParam(r, "where")
		fmt.Println(where)
		url := "https://catastro.crimpr.net/proxy/proxy.ashx?https://catastro.crimpr.net/server/rest/services/Parcelario/Parcelas/MapServer/654/query?f=json&returnIdsOnly=true&returnCountOnly=true&where="+ where + "&returnGeometry=false&spatialRel=esriSpatialRelIntersects&outSR=102100"
		fmt.Println(url)
		res, err := http.Get(url)
		if err != nil{
			panic(err)
		}
		bA, err := io.ReadAll(res.Body)
		if err != nil{
			panic(err)
		}
		w.Write(bA)
	})

	http.ListenAndServe(":3000", r)
}

