package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	helpers "github.com/nachoconques0/diagnosis_svc/internal/helpers/db"
	userRepo "github.com/nachoconques0/diagnosis_svc/internal/repo/user"
)

func insertTestUser(t *testing.T, db *gorm.DB) *user.Entity {
	u := &user.Entity{
		ID:        uuid.New(),
		Nickname:  "nacho",
		Password:  "doctordoctor:)",
		Email:     "nacho@test.com",
		CreatedAt: time.Now(),
	}
	err := db.Create(u).Error
	assert.NoError(t, err)
	return u
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db, teardown, err := helpers.NewTestDB()
	if err != nil {
		assert.Nil(t, err)
	}
	defer teardown()
	repo := userRepo.NewRepository(db)
	ctx := context.Background()

	createdUser := insertTestUser(t, db)

	t.Run("user found", func(t *testing.T) {
		u, err := repo.GetByEmail(ctx, createdUser.Email)
		assert.NoError(t, err)
		assert.Equal(t, createdUser.Email, u.Email)
		assert.Equal(t, createdUser.Nickname, u.Nickname)
	})

	t.Run("user not found", func(t *testing.T) {
		u, err := repo.GetByEmail(ctx, "notfound@test.com")
		assert.Error(t, err)
		assert.Nil(t, u)
	})
}
