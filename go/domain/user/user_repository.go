package user

type UserRepository interface {
	FindById(id string) (*User, error)
}
