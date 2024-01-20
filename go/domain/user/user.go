package user

import "github.com/google/uuid"

type User struct {
	ID    string
	Email string
	Type  string
}

func NewUser(email string, userType string) *User {
	return &User{
		ID:    uuid.New().String(),
		Email: email,
		Type:  userType,
	}
}
