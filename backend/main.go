package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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
	r.Post("/", func (w http.ResponseWriter, r *http.Request) {
		body := r.Body
		var prop Property
		err := json.NewDecoder(body).Decode(&prop)
		if err != nil{
			panic(err)
		}

		where := prop.createUrl()
		url := "https://catastro.crimpr.net/proxy/proxy.ashx?https://catastro.crimpr.net/server/rest/services/Parcelario/Parcelas/MapServer/654/query?f=json&where=" + url.PathEscape(where) + "&returnGeometry=true&spatialRel=esriSpatialRelIntersects&outFields=OLDPID,NUM_CATASTRO,TIPO,CATASTRO,MUNICIPIO,CONTACT,DIRECCION_FISICA,DIRECCION_POSTAL,CABIDA,LAND,STRUCTURE,MACHINERY,TOTALVAL,EXEMP,EXON,TAXABLE,DEEDBOOK,DEEDPAGE,ESTATE,DEEDNUM,SALESAMT,SALESDTTM,SELLERNAME,BYERNAME,Shape.STArea(),Shape.STLength(),OBJECTID_1,INSIDE_X,INSIDE_Y,OBJECTID&outSR=102100&resultOffset=0&resultRecordCount=2000"
		fmt.Println(url)
		res, err := http.Get(url)
		if err != nil{
			panic(err)
		}
		response, err := io.ReadAll(res.Body)
		if err != nil{
			panic(err)
		}
		w.Write(response)
	})

	http.ListenAndServe(":3000", r)
}

func (p *Property) createUrl() string {
	filters := []string{}

	if p.FirstName != ""{
		filters = append(filters, fmt.Sprintf("CONTACT LIKE '%s'", "%" + p.FirstName + "%"))
	}


	if p.LastName != ""{
		filters = append(filters, fmt.Sprintf("CONTACT LIKE '%s'", "%" + p.LastName + "%"))
	}


	if p.Catastro != ""{
		filters = append(filters, fmt.Sprintf("CATASTRO LIKE '%s'", "%" + p.Catastro + "%"))
	}


	if p.DireccionFisica != ""{
		filters = append(filters, fmt.Sprintf("DIRECCION_FISICA LIKE '%s'", "%" + p.DireccionFisica + "%"))
	}


	if p.DireccionPostal != ""{
		filters = append(filters, fmt.Sprintf("DIRECCION_POSTAL LIKE '%s'", "%" + p.DireccionPostal + "%"))
	}

	s := strings.Join(filters, " AND ")

	return s
}
