package db

import (
	"api1-new/pkg/logger"
)

func GetCompanies() ([]company, error) {
	logger.Log.Info("SELECT id, name FROM company;")
	// sqlStatement := `SELECT id, name FROM company`
	// rows, err := RunSQL(sqlStatement)
	// if err != nil {
	// 	return nil, err
	// }
	// var companies []company
	// for rows.Next() {
	// 	var row company
	// 	err := rows.Scan(&row.Id, &row.Name)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	companies = append(companies, row)
	// }

	return mockCompanies, nil
}

func CreateCompany(name string) error {
	logger.Log.Info("INSERT INTO company (name) VALUES ($1); " + name)
	// sqlStatement := `INSERT INTO company (name) VALUES ($1)`
	// _, err := RunSQL(sqlStatement, name)
	// if err != nil {
	// 	return err
	// }
	mockCompanies = append(mockCompanies, company{Name: name, Id: len(mockCompanies) + 1})
	return nil
}

type company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var mockCompanies = []company{
	{Id: 1, Name: "Company 1"},
	{Id: 2, Name: "Company 2"},
	{Id: 3, Name: "Company 3"},
}
