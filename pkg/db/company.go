package db

import (
	"api1-new/pkg/logger"
)

func GetCompanies() []company {
	logger.Log.Info(runSql("SELECT id, name FROM company"))
	return []company{
		{Id: 1, Name: "Company 1"},
		{Id: 2, Name: "Company 2"},
		{Id: 3, Name: "Company 3"},
	}
}

type company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
