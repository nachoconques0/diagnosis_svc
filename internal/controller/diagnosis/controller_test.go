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

	controller "github.com/nachoconques0/diagnosis_svc/internal/controller/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/mocks"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateDiagnosis(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDiagnosisService(ctrl)
	controller := controller.New(mockService)

	prescription := "rest"
	req := model.CreateDiagnosisRequest{
		PatientID:    "1234",
		Diagnosis:    "gripessilla",
		Prescription: &prescription,
	}
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		jsonBody, _ := json.Marshal(map[string]string{
			"patient_id":   req.PatientID,
			"diagnosis":    req.Diagnosis,
			"prescription": *req.Prescription,
		})

		c.Request, _ = http.NewRequest(http.MethodPost, "/diagnosis", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		expected := &model.DiagnosisResponse{
			Diagnosis: req.Diagnosis,
		}
		mockService.EXPECT().
			Create(gomock.Any(), req).
			Return(expected, nil)

		r.POST("/diagnosis", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("invalid req", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		c.Request, _ = http.NewRequest(http.MethodPost, "/diagnosis", bytes.NewBuffer([]byte(`{}`)))
		c.Request.Header.Set("Content-Type", "application/json")

		r.POST("/diagnosis", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		jsonBody, _ := json.Marshal(map[string]string{
			"patient_id":   req.PatientID,
			"diagnosis":    req.Diagnosis,
			"prescription": *req.Prescription,
		})

		c.Request, _ = http.NewRequest(http.MethodPost, "/diagnosis", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		mockService.EXPECT().
			Create(gomock.Any(), req).
			Return(nil, errors.New("unexpected error"))

		r.POST("/diagnosis", controller.Create)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestFindDiagnosis(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDiagnosisService(ctrl)
	controller := controller.New(mockService)

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		req, _ := http.NewRequest(http.MethodGet, "/diagnosis?name=nachin", nil)
		c.Request = req

		mockService.EXPECT().
			Find(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]model.DiagnosisResponse{{Diagnosis: "lagripe"}}, nil)

		r.GET("/diagnosis", controller.Find)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, r := gin.CreateTestContext(rec)

		req, _ := http.NewRequest(http.MethodGet, "/diagnosis?name=nachin", nil)
		c.Request = req

		mockService.EXPECT().
			Find(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errors.New("someERr"))

		r.GET("/diagnosis", controller.Find)
		r.ServeHTTP(rec, c.Request)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
