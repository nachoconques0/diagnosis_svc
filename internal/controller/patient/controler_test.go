package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	controller "github.com/nachoconques0/diagnosis_svc/internal/controller/patient"
	patientEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPatientService(ctrl)
	controller := controller.New(mockService)

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		jsonBody, _ := json.Marshal(map[string]string{
			"name":  "nacho",
			"email": "jcalcagno@gmail.com",
			"dni":   "123123",
		})

		c.Request, _ = http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		expected := &patientEntity.Entity{
			Name: "nacho",
		}
		mockService.EXPECT().
			Create(gomock.Any(), "nacho", "jcalcagno@gmail.com", "123123", gomock.Any(), gomock.Any()).
			Return(expected, nil)

		r.POST("/patients", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		jsonBody, _ := json.Marshal(map[string]string{
			"name":  "nacho",
			"email": "jcalcagno@gmail.com",
			"dni":   "123123",
		})

		c.Request, _ = http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		mockService.EXPECT().
			Create(gomock.Any(), "nacho", "jcalcagno@gmail.com", "123123", gomock.Any(), gomock.Any()).
			Return(nil, errors.New("unexpected error"))

		r.POST("/patients", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestFindDiagnosis(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPatientService(ctrl)
	controller := controller.New(mockService)

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		req, _ := http.NewRequest(http.MethodGet, "/patients?name=nachin", nil)
		c.Request = req

		mockService.EXPECT().
			Find(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]patientEntity.Entity{{Name: "nachin"}}, nil)

		r.GET("/diagnosis", controller.Find)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		req, _ := http.NewRequest(http.MethodGet, "/patients?name=nachin", nil)
		c.Request = req

		mockService.EXPECT().
			Find(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errors.New("someERr"))

		r.GET("/patients", controller.Find)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
