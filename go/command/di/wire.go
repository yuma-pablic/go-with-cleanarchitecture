//go:build wireinject
// +build wireinject

// wire.go
package di

import (
	"profile/infra/rdb/repository"
	userPresenter "profile/presentation/user"
	userUseCase "profile/usecase/user"

	"github.com/google/wire"
)

func InitUser() userPresenter.Handler {
	wire.Build(repository.NewUserRepository, userUseCase.NewFindUserUseCase, userPresenter.NewHandler)
	return userPresenter.Handler{}
}
