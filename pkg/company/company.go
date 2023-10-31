package company

import "api1-new/pkg/db"

func GetCompany() ([]Company, error) {
	companiesDB, err := db.GetCompanies()
	if err != nil {
		return nil, err
	}
	var companies []Company
	for _, companyDB := range companiesDB {
		company := Company{
			Id:   companyDB.Id,
			Name: companyDB.Name,
		}
		companies = append(companies, company)
	}
	return companies, nil
}

func CreateCompany(name string) error {
	err := db.CreateCompany(name)
	if err != nil {
		return err
	}
	return nil
}

type Company struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}
