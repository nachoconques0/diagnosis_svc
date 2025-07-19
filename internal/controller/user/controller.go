package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type service interface {
	GetByEmail(ctx context.Context, email string) (*user.Entity, error)
}

// Controller holds the required dependencies in order to implement the logic service of the user requests
type Controller struct {
	svc    service
	jwtKey string
}

// New returns a new HTTP Controller with the given service implementation
func New(svc service, jwtKey string) *Controller {
	return &Controller{svc: svc, jwtKey: jwtKey}
}

// Login will check if user exists and create a JWT token for auth
func (ctrl *Controller) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("loging failed when ShouldBindJSON err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.svc.GetByEmail(c.Request.Context(), req.Email)
	if err != nil {
		if appErr, ok := err.(*errors.Error); ok {
			log.Error().Err(err).Msg(fmt.Sprintf("loging failed in GetByEmail with err: %s", err.Error()))
			c.JSON(appErr.HTTPStatus(), appErr)
			return
		}
		log.Error().Err(err).Msg(fmt.Sprintf("loging failed in GetByEmail with err: %s", err.Error()))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := comparePasswords(res.Password, req.Password); err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("loging failed checking pw with err: %s", err.Error()))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, err := createToken(*res, ctrl.jwtKey)
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("loging failed creating token with err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginResponse{Token: tokenString})
}

func createToken(u user.Entity, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"nickname": u.Nickname,
			"email":    u.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func comparePasswords(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}
