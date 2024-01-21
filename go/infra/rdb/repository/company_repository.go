package repository

import companyDomain "profile/domain/company"

type companyRepository struct{}

func NewCompanyRepository() companyDomain.CompanyRepository {
	return &companyRepository{}
}

func (cr *companyRepository) Find() (*companyDomain.Company, error) {
	return &companyDomain.Company{
		ID:                "1234567810",
		Name:              "Example Company",
		NumberOfEmployees: 100,
	}, nil
}
