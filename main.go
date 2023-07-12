package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	apiGroup := server.Group("/api")
	personGroup := apiGroup.Group("/person")
	{
		personGroup.GET("", GetPersonEndpoint)
		personGroup.POST("", PostPersonEndpoint)
	}
	companyGroup := apiGroup.Group("/company")
	{
		companyGroup.GET("", GetCompanyEndpoint)
	}
	server.Run("0.0.0.0:8080")
}

// until now in /endpoints/personEndpoint.go
func GetPersonEndpoint(c *gin.Context) {
	data := GetPersonService()
	c.JSON(200, data)
}

// until now in /endpoints/personEndpoint.go
func PostPersonEndpoint(c *gin.Context) {
	// placeholder
	c.JSON(200, "Placeholder")
}

// until now in /endpoints/companiesEndpoint.go
func GetCompanyEndpoint(c *gin.Context) {
	data := GetCompanyService()
	c.JSON(200, data)
}

// until now in /services/personService.go
func GetPersonService() []Person {
	persons := []Person{}
	personDBs := GetPersonDao()
	companies := GetCompaniesDao()
	for _, personDB := range personDBs {
		person := Person{
			Name: personDB.Name,
			Age:  personDB.Age,
		}
		for _, company := range companies {
			if company.Id == personDB.CompanyId {
				person.CompanyName = company.Name
				break
			}
		}
		persons = append(persons, person)
	}
	return persons
}

// until now in /services/companiesService.go
func GetCompanyService() []Company {
	companies := GetCompaniesDao()
	return companies
}

// until now in /dao/personDao.go
func GetPersonDao() []PersonDB {
	Log.Info(RunSql("SELECT name, age, companyid FROM person"))
	return []PersonDB{
		{Name: "Person 1", Age: 20, CompanyId: 1},
		{Name: "Person 2", Age: 30, CompanyId: 2},
		{Name: "Person 3", Age: 40, CompanyId: 3},
		{Name: "Person 4", Age: 50, CompanyId: 2},
	}
}

// until now in /dao/companiesDao.go
func GetCompaniesDao() []Company {
	Log.Info(RunSql("SELECT id, name FROM company"))
	return []Company{
		{Id: 1, Name: "Company 1"},
		{Id: 2, Name: "Company 2"},
		{Id: 3, Name: "Company 3"},
	}
}

// until now in /utils/db.go
func RunSql(sql string) string {
	// Placeholder for make DB select
	return sql
}

// until now in /utils/logger.go
var (
	Log *zap.Logger
)

// until now in /utils/logger.go
func InitializeZapCustomLogger() {
	loglevel := zapcore.DebugLevel
	conf := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(loglevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			MessageKey:   "msg",
			CallerKey:    "file",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	foo, err := conf.Build()
	if err != nil {
		Log.Error(err.Error())
	}
	Log = foo
}

type PersonDB struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CompanyId int    `json:"companyId"`
}

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	CompanyName string `json:"company"`
}

type Company struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
