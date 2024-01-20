package user

type User struct {
	ID    string
	Email string
	Type  string
}

func NewUser(id string, email string, userType string) *User {
	return &User{id, email, userType}
}
