package user

import (
	"context"
	userDomain "profile/domain/user"
)

type CreateUserUseCase struct {
	ur userDomain.UserRepository
}

type CreateUserUseCaseInputDTO struct {
	ID    string
	Email string
	Type  string
}

func NewCreateUserUseCase(ur userDomain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{ur}
}

func (uc *CreateUserUseCase) Run(ctx context.Context, input *CreateUserUseCaseInputDTO) error {
	user := userDomain.NewUser(input.Email, input.Type)
	err := uc.ur.Create(user)
	if err != nil {
		return err
	}
	return nil
}
