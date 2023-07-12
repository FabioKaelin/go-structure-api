package person

import "api1-new/pkg/db"

func GetPerson() []Person {
	persons := []Person{}
	personDBs := db.GetPerson()
	companies := db.GetCompanies()
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

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	CompanyName string `json:"company"`
}
