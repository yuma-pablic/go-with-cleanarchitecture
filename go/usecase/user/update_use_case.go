package user

import (
	userDomain "profile/domain/user"
)

type UpdateUserUseCase struct {
	ur userDomain.UserRepository
	us userDomain.UserDomainService
}

type UpdateUserUseCaseInputDTO struct {
	ID    string
	Email string
	Type  string
}

func NewUpdateUserUseCase(ur userDomain.UserRepository, us userDomain.UserDomainService) *UpdateUserUseCase {
	return &UpdateUserUseCase{ur, us}
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
	err = uc.ur.Update(user)
	if err != nil {
		return err
	}
	return nil
}
