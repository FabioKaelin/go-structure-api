package person

import "api1-new/pkg/db"

func GetPerson() ([]Person, error) {
	persons := []Person{}
	personDBs := db.GetPerson()
	companies, err := db.GetCompanies()
	if err != nil {
		return nil, err
	}

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
	return persons, nil
}

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	CompanyName string `json:"company"`
}
