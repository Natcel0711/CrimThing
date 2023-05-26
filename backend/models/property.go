package models

import(
	"fmt"
	"strings"
)

type Property struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Catastro string `json:"catastro"`
	DireccionFisica string `json:"direccion_fisica"`
	DireccionPostal string `json:"direccion_postal"`
}


func (p *Property) CreateUrl() string {
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

	if len(filters) == 0{
		return ""
	}

	s := strings.Join(filters, " AND ")

	return s
}
