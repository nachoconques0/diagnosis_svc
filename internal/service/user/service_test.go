package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	userEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
	userService "github.com/nachoconques0/diagnosis_svc/internal/service/user"
)

func TestUserService_GetByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockUserRepository(ctrl)
	svc := userService.New(repo)
	ctx := context.Background()
	email := "nachintimetosleep@gmail.com"
	password := "123123"

	t.Run("user found", func(t *testing.T) {
		repo.EXPECT().GetByEmail(ctx, email).Return(&userEntity.Entity{Email: email, Password: password}, nil)
		res, err := svc.GetByEmail(ctx, email, password)
		assert.NoError(t, err)
		assert.Equal(t, email, res.Email)
	})

	t.Run("user not found", func(t *testing.T) {
		repo.EXPECT().GetByEmail(ctx, email).Return(nil, errors.New("not found"))
		res, err := svc.GetByEmail(ctx, email, password)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
