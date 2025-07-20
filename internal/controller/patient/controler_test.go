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
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPatientService(ctrl)
	controller := controller.New(mockService)

	patientTestName := "nachin"
	patientTestDNI := "123123"
	patientTestEmail := "nacho@gmail.com"

	req := model.CreatePatientRequest{
		Name:  patientTestName,
		DNI:   patientTestDNI,
		Email: patientTestEmail,
	}

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		jsonBody, _ := json.Marshal(map[string]string{
			"name":  patientTestName,
			"email": patientTestEmail,
			"dni":   patientTestDNI,
		})

		c.Request, _ = http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		expected := &model.PatientResponse{
			Name: patientTestName,
		}
		mockService.EXPECT().
			Create(gomock.Any(), req).
			Return(expected, nil)

		r.POST("/patients", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		jsonBody, _ := json.Marshal(map[string]string{
			"name":  patientTestName,
			"email": patientTestEmail,
			"dni":   patientTestDNI,
		})

		c.Request, _ = http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		mockService.EXPECT().
			Create(gomock.Any(), req).
			Return(nil, errors.New("unexpected error"))

		r.POST("/patients", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestFindPatients(t *testing.T) {
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
			Return([]model.PatientResponse{{Name: "nachin"}}, nil)

		r.GET("/patients", controller.Find)
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
