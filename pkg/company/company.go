package company

import "api1-new/pkg/db"

func GetCompany() ([]company, error) {
	companiesDB, err := db.GetCompanies()
	if err != nil {
		return nil, err
	}
	var companies []company
	for _, companyDB := range companiesDB {
		company := company{
			Id:   companyDB.Id,
			Name: companyDB.Name,
		}
		companies = append(companies, company)
	}
	return companies, nil
}

type company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
