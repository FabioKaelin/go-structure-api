package db

import (
	"api1-new/pkg/config"
	"api1-new/pkg/logger"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

func runSql(sql string) string {
	// Placeholder for make DB select
	return sql
}

var dbConn *sql.DB

func UpdateDBConnection() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	dbNew, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		logger.Log.Error(err.Error())
	}
	// test if connection is working
	err = dbNew.Ping()
	if err != nil {
		newErr := errors.Join(errors.New("error durring updating db connection"), err)
		return newErr
	}
	dbConn = dbNew
	// db.QueryRow("set client_encoding='win1252'")
	// db.QueryRow("SET CLIENT_ENCODING TO 'LATIN1';")
	return nil
}

func RunSQL(sqlStatement string, parameters ...any) (*sql.Rows, error) {

	err := dbConn.Ping()
	if err != nil {
		logger.Log.Warn("DB Connection lost, reconnecting...")
		err := UpdateDBConnection()
		if err != nil {
			return &sql.Rows{}, err
		}
	}
	rows, err := dbConn.Query(sqlStatement, parameters...)
	if err != nil {
		newErr := errors.Join(errors.New("error durring executing "+sqlStatement), err)
		return &sql.Rows{}, newErr
	}
	return rows, nil
}
