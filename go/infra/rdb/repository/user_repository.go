package repository

import (
	userDomain "profile/domain/user"
)

type userRepository struct{}

func NewUserRepository() userDomain.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) FindById(id string) (*userDomain.User, error) {
	return &userDomain.User{
		ID:    "1234567810",
		Email: "example@gmail.com",
		Type:  "employee",
	}, nil
}

func (ur *userRepository) Create(*userDomain.User) error {
	return nil
}

func (ur *userRepository) Update(*userDomain.User) error {
	return nil
}
