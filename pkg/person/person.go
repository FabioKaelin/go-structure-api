package person

import "api1-new/pkg/db"

func GetPerson() ([]Person, error) {
	persons := []Person{}
	personDBs, err := db.GetPerson()
	if err != nil {
		return nil, err
	}
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

func CreatePerson(name string, age int, companyId int) error {
	err := db.CreatePerson(name, age, companyId)
	if err != nil {
		return err
	}
	return nil
}

type Person struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	CompanyName string `json:"company"`
}

type PersonPost struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CompanyId int    `json:"companyId"`
}
