package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
	"github.com/rs/zerolog/log"
)

type service interface {
	Create(ctx context.Context, req model.CreateDiagnosisRequest) (*model.DiagnosisResponse, error)
	Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]model.DiagnosisResponse, error)
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
	var input model.CreateDiagnosisRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("Create failed when ShouldBindJSON err: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	res, err := ctrl.svc.Create(c.Request.Context(), input)
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
		PatientName: strings.Trim(patientName, `"`),
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
