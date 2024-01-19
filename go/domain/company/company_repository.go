package company

type CompanyRepository interface {
	FindById(id string) (*Company, error)
}
