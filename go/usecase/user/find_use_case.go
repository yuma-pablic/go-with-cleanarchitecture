package user

import (
	"context"
	userDomain "profile/domain/user"
)

type FindUserUseCase struct {
	ur userDomain.UserRepository
}

type FindUserUseCaseOutputDTO struct {
	ID    string
	Email string
	Type  string
}

func NewFindUserUseCase(ur userDomain.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{ur}
}

func (uc *FindUserUseCase) Run(ctx context.Context, id string) (*FindUserUseCaseOutputDTO, error) {
	user, err := uc.ur.FindById(id)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseOutputDTO{
		ID:    user.ID,
		Email: user.Email,
		Type:  user.Type,
	}, nil
}
