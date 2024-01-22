package company

type Company struct {
	ID                string
	Name              string
	NumberOfEmployees int
}

func ReNewCompany(id string, name string, numberOfEmployees int) *Company {
	return &Company{
		ID:                id,
		Name:              name,
		NumberOfEmployees: numberOfEmployees,
	}
}
