package user

import (
	"context"
	userDomain "profile/domain/user"
)

type CreateUserUseCase struct {
	ur  userDomain.UserRepository
	uds userDomain.UserDomainService
}

type CreateUserUseCaseInputDTO struct {
	Email string
	Type  string
}

func NewCreateUserUseCase(ur userDomain.UserRepository, uds userDomain.UserDomainService) *CreateUserUseCase {
	return &CreateUserUseCase{ur, uds}
}

func (uc *CreateUserUseCase) Run(ctx context.Context, input *CreateUserUseCaseInputDTO) error {
	user, err := uc.uds.NewUser(input.Email, input.Type)
	if err != nil {
		return err
	}
	err = uc.ur.Create(user)
	if err != nil {
		return err
	}
	return nil
}
