package company

import "api1-new/pkg/db"

func GetCompany() []company {
	companiesDB := db.GetCompanies()
	var companies []company
	for _, companyDB := range companiesDB {
		company := company{
			Id:   companyDB.Id,
			Name: companyDB.Name,
		}
		companies = append(companies, company)
	}
	return companies
}

type company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
