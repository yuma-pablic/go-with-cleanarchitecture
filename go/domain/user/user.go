package user

import (
	"errors"
	"profile/domain/company"

	"github.com/google/uuid"
)

type User struct {
	ID    string
	Email string
	Type  string
}
type UserDomainService struct {
	ur UserRepository
	cr company.CompanyRepository
}

func NewUserDomainService(ur UserRepository, cr company.CompanyRepository) *UserDomainService {
	return &UserDomainService{ur, cr}
}
func (uds *UserDomainService) NewUser(email string, userType string) (*User, error) {
	emailDomain := email
	company, err := uds.cr.Find()
	if err != nil {
		return nil, errors.New("company not found")
	}
	if emailDomain == company.Name {
		userType = "company"
	} else {
		userType = "general"
	}
	return newUser(uuid.New().String(), email, userType), nil
}

func (uds *UserDomainService) ReNewUser(id string, email string, userType string) (*User, error) {
	user, err := uds.ur.FindById(id)
	if err != nil {
		return nil, err
	}
	if user.Email != email {
		return nil, errors.New("email is not match")
	}
	emailDomain := user.Email
	company, err := uds.cr.Find()
	if err != nil {
		return nil, err
	}
	if emailDomain != company.Name {
		userType = "general"
	} else {
		userType = "company"
	}
	return newUser(id, email, userType), nil
}

// プライマリーコンストラクタ
func newUser(id string, email string, userType string) *User {
	return &User{
		ID:    id,
		Email: email,
		Type:  userType,
	}
}
