package user

import (
	"profile/domain/company"
	userDomain "profile/domain/user"
)

type UpdateUserUseCase struct {
	ur userDomain.UserRepository
	us userDomain.UserDomainService
	cr company.CompanyRepository
}

type UpdateUserUseCaseInputDTO struct {
	ID    string
	Email string
	Type  string
}

func NewUpdateUserUseCase(ur userDomain.UserRepository, us userDomain.UserDomainService, cr company.CompanyRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{ur, us, cr}
}

func (uc *UpdateUserUseCase) Run(input *UpdateUserUseCaseInputDTO) error {
	response, err := uc.ur.FindById(input.ID)
	if err != nil {
		return err
	}
	user, err := uc.us.ReNewUser(response.ID, input.Email, input.Type)
	if err != nil {
		return err
	}
	res, err := uc.cr.Find()
	if err != nil {
		return err
	}
	company := company.ReNewCompany(res.ID, res.Name, res.NumberOfEmployees)

	err = uc.ur.Update(user)
	err = uc.cr.Update(company)
	if err != nil {
		return err
	}
	return nil
}
