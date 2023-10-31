package db

import (
	"api1-new/pkg/logger"
	"fmt"
)

func GetPerson() ([]person, error) {
	logger.Log.Info("SELECT name, age, companyid FROM person")
	// sqlStatement := `SELECT name, age, companyid FROM person`
	// rows, err := RunSQL(sqlStatement)
	// if err != nil {
	// 	return nil, err
	// }
	// var persons []person
	// for rows.Next() {
	// 	var row person
	// 	err := rows.Scan(&row.Name, &row.Age, &row.CompanyId)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	persons = append(persons, row)
	// }
	return mockPersons, nil
}

func CreatePerson(name string, age int, companyId int) error {
	logger.Log.Info("INSERT INTO person (name, age, companyid) VALUES ($1, $2, $3); " + name + ", " + fmt.Sprint(age) + ", " + fmt.Sprint(companyId))
	// sqlStatement := `INSERT INTO person (name, age, companyid) VALUES ($1, $2, $3)`
	// _, err := RunSQL(sqlStatement, name, age, companyId)
	// if err != nil {
	// 	return err
	// }
	mockPersons = append(mockPersons, person{Name: name, Age: age, CompanyId: companyId})
	return nil
}

type person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CompanyId int    `json:"companyId"`
}

var mockPersons = []person{
	{Name: "Person 1", Age: 20, CompanyId: 1},
	{Name: "Person 2", Age: 30, CompanyId: 2},
	{Name: "Person 3", Age: 40, CompanyId: 3},
	{Name: "Person 4", Age: 50, CompanyId: 2},
}
