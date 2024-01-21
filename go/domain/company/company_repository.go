package company

type CompanyRepository interface {
	Find() (*Company, error)
}
