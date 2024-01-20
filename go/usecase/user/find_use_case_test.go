package user

import (
	"context"
	"os/user"
	"testing"

	mocks "profile/domain/user/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFindUseCa(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	ur := NewFindUserUseCase(mockUserRepo)
	id := "aaa"
	tests := []struct {
		name     string
		mockFunc func()
		want     []*FindUserUseCaseOutputDTO
		wantErr  bool
	}{
		{
			name: "success,get_entries",
			mockFunc: func() {
				mockUserRepo.EXPECT().FindById(id).Return(&user.User{}, nil)
			},
			want: []*FindUserUseCaseOutputDTO{
				{
					ID:    id,
					Email: "",
					Type:  "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			tt.mockFunc()
			response, err := ur.Run(context.Background(), id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, response)
		})
	}
}
