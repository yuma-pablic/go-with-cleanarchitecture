package user

import "github.com/google/uuid"

type User struct {
	ID    string
	Email string
	Type  string
}

func NewUser(email string, userType string) *User {
	return newUser(uuid.New().String(), email, userType)
}

func newUser(id string, email string, userType string) *User {
	return &User{
		ID:    id,
		Email: email,
		Type:  userType,
	}
}
