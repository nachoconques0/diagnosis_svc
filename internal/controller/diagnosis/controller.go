package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	diagEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/rs/zerolog/log"
)

type service interface {
	Create(ctx context.Context, patientID string, diag string, prescription *string) (*diagEntity.Entity, error)
	Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]diagEntity.Entity, error)
}

// Controller holds the required dependencies in order to implement the logic service of the diagnosis requests
type Controller struct {
	svc service
}

// New returns a new HTTP Controller with the given service implementation
func New(svc service) *Controller {
	return &Controller{svc: svc}
}

// Create creates a diagnosis for an specific patient
func (ctrl *Controller) Create(c *gin.Context) {
	var input createDiagnosisInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("Create failed when ShouldBindJSON err: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	res, err := ctrl.svc.Create(c.Request.Context(), input.PatientID, input.Diagnosis, input.Prescription)
	if err != nil {
		if appErr, ok := err.(*errors.Error); ok {
			log.Error().Err(err).Msg(fmt.Sprintf("Create failed with err: %s", err.Error()))
			c.JSON(appErr.HTTPStatus(), appErr)
			return
		}
		log.Error().Err(err).Msg(fmt.Sprintf("Create failed with err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// Find returns a list of diagnoses based on pagination, patient name & date
func (ctrl *Controller) Find(c *gin.Context) {
	patientName := c.Query("name")
	parsedDate := c.Query("date")

	var date *time.Time
	if parsedDate != "" {
		parsed, err := time.Parse("2006-01-02", parsedDate)
		if err == nil {
			date = &parsed
		}
	}

	pagination := query.NewPagination(c.Query("page"), c.Query("page_size"))

	res, err := ctrl.svc.Find(c.Request.Context(), query.DiagnosisFilters{
		PatientName: patientName,
		Date:        date,
	}, pagination)
	if err != nil {
		if appErr, ok := err.(*errors.Error); ok {
			log.Error().Err(err).Msg(fmt.Sprintf("Find failed with err: %s", err.Error()))
			c.JSON(appErr.HTTPStatus(), appErr)
			return
		}
		log.Error().Err(err).Msg(fmt.Sprintf("Find failed with err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

type createDiagnosisInput struct {
	PatientID    string  `json:"patient_id" binding:"required"`
	Diagnosis    string  `json:"diagnosis" binding:"required"`
	Prescription *string `json:"prescription,omitempty"`
}
