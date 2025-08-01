package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
)

type service interface {
	Create(ctx context.Context, req model.CreatePatientRequest) (*model.PatientResponse, error)
	Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]model.PatientResponse, error)
}

// Controller holds the required dependencies in order to implement the logic service of the patient requests
type Controller struct {
	svc service
}

// New returns a new HTTP Controller with the given service implementation
func New(svc service) *Controller {
	return &Controller{svc: svc}
}

// Create creates a diagnosis for an specific patient
func (ctrl *Controller) Create(c *gin.Context) {
	var input model.CreatePatientRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("Create failed when ShouldBindJSON err: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	res, err := ctrl.svc.Create(c.Request.Context(), input)
	if err != nil {
		if appErr, ok := err.(*errors.Error); ok {
			log.Error().Err(err).Msg(fmt.Sprintf("Create patient failed with err: %s", err.Error()))
			c.JSON(appErr.HTTPStatus(), appErr)
			return
		}
		log.Error().Err(err).Msg(fmt.Sprintf("Create patient failed with err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// Find returns a list of patients
func (ctrl *Controller) Find(c *gin.Context) {
	patientName := c.Query("name")

	pagination := query.NewPagination(c.Query("page"), c.Query("page_size"))

	res, err := ctrl.svc.Find(c.Request.Context(), query.DiagnosisFilters{
		PatientName: patientName,
	}, pagination)
	if err != nil {
		if appErr, ok := err.(*errors.Error); ok {
			log.Error().Err(err).Msg(fmt.Sprintf("Find patients failed with err: %s", err.Error()))
			c.JSON(appErr.HTTPStatus(), appErr)
			return
		}
		log.Error().Err(err).Msg(fmt.Sprintf("Find patients failed with err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
