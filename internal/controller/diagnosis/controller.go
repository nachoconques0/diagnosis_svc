package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	diagEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
)

type service interface {
	Create(ctx context.Context, patientID string, diag string, prescription *string) (*diagEntity.Entity, error)
	Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]diagEntity.Entity, error)
}

type Controller struct {
	svc service
}

func New(svc service) *Controller {
	return &Controller{svc: svc}
}

type createDiagnosisInput struct {
	PatientID    string  `json:"patient_id" binding:"required"`
	Diagnosis    string  `json:"diagnosis" binding:"required"`
	Prescription *string `json:"prescription,omitempty"`
}

func (ctrl *Controller) Create(c *gin.Context) {
	var input createDiagnosisInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	res, err := ctrl.svc.Create(c.Request.Context(), input.PatientID, input.Diagnosis, input.Prescription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (d *Controller) Find(c *gin.Context) {
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

	res, err := d.svc.Find(c.Request.Context(), query.DiagnosisFilters{
		PatientName: patientName,
		Date:        date,
	}, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
