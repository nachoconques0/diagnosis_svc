package server

import (
	"github.com/gin-gonic/gin"

	diagnosisCtrl "github.com/nachoconques0/diagnosis_svc/internal/controller/diagnosis"
	patientCtrl "github.com/nachoconques0/diagnosis_svc/internal/controller/patient"
	userCtrl "github.com/nachoconques0/diagnosis_svc/internal/controller/user"
	"github.com/nachoconques0/diagnosis_svc/internal/http/middleware"
)

// InitRoutes will set all the endpoints needed for the service
func InitRoutes(
	router *gin.Engine,
	userCtrl userCtrl.Controller,
	diagnosisCtrl diagnosisCtrl.Controller,
	patientCtrl patientCtrl.Controller,
) {
	v1 := router.Group("/v1")

	v1.POST("/login", userCtrl.Login)

	v1.POST("/diagnosis", middleware.ProtectedHandler(), diagnosisCtrl.Create)
	v1.GET("/diagnosis", middleware.ProtectedHandler(), diagnosisCtrl.Find)

	v1.POST("/patients", middleware.ProtectedHandler(), patientCtrl.Create)
	v1.GET("/patients", middleware.ProtectedHandler(), patientCtrl.Find)
}
