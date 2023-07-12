package db

import "api1-new/pkg/logger"

func GetPerson() []person {
	logger.Log.Info(runSql("SELECT name, age, companyid FROM person"))
	return []person{
		{Name: "Person 1", Age: 20, CompanyId: 1},
		{Name: "Person 2", Age: 30, CompanyId: 2},
		{Name: "Person 3", Age: 40, CompanyId: 3},
		{Name: "Person 4", Age: 50, CompanyId: 2},
	}
}

type person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CompanyId int    `json:"companyId"`
}
