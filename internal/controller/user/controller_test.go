package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"

	controller "github.com/nachoconques0/diagnosis_svc/internal/controller/user"
	userEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
)

func TestLoginController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mocks.NewMockUserService(ctrl)
	jwtKey := "testsecret"
	c := controller.New(svc, jwtKey)
	testEmail := "nachin@gmail.com"

	t.Run("success", func(t *testing.T) {
		password := "testpass"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		svc.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(&userEntity.Entity{
			Email:    testEmail,
			Password: string(hash),
			Nickname: "nacho",
		}, nil)

		w := httptest.NewRecorder()
		r := gin.New()
		r.POST("/login", c.Login)

		body, _ := json.Marshal(map[string]string{
			"email":    testEmail,
			"password": password,
		})
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "token")
	})

	t.Run("invalid password", func(t *testing.T) {
		password := "testpass"
		hash, _ := bcrypt.GenerateFromPassword([]byte("diferenteee"), bcrypt.DefaultCost)
		svc.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(&userEntity.Entity{
			Email:    testEmail,
			Password: string(hash),
			Nickname: "nacho",
		}, nil)

		w := httptest.NewRecorder()
		r := gin.New()
		r.POST("/login", c.Login)

		body, _ := json.Marshal(map[string]string{
			"email":    testEmail,
			"password": password,
		})
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid credentials")
	})

	t.Run("service error", func(t *testing.T) {
		password := "any"
		svc.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, errors.New("userNotFound"))

		w := httptest.NewRecorder()
		r := gin.New()
		r.POST("/login", c.Login)

		body, _ := json.Marshal(map[string]string{
			"email":    testEmail,
			"password": password,
		})
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "userNotFound")
	})
}
