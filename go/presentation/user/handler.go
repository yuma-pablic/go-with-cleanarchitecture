package user

import (
	"context"
	"net/http"
	"profile/usecase/user"
)

type Handler struct {
	FindUserUseCase *user.FindUserUseCase
}

func NewHandler(fu *user.FindUserUseCase) Handler {
	return Handler{
		FindUserUseCase: fu,
	}
}

func (h Handler) GetByUserID(ctx context.Context, r *http.Request) (*userResponseModel, error) {
	id := r.URL.Query().Get("id")
	dtos, err := h.FindUserUseCase.Run(ctx, id)
	if err != nil {
		return nil, err
	}
	return &userResponseModel{
		ID:    dtos.ID,
		Email: dtos.Email,
		Type:  dtos.Type,
	}, nil

}
