package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

type service interface {
	GetByEmail(ctx context.Context, email, password string) (*user.Entity, error)
}

type Controller struct {
	svc    service
	jwtKey string
}

func New(svc service, jwtKey string) *Controller {
	return &Controller{svc: svc, jwtKey: jwtKey}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (ctrl *Controller) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if customErr, ok := err.(*errors.Error); ok {
			c.JSON(customErr.HTTPStatus(), customErr)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.svc.GetByEmail(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := comparePasswords(res.Password, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, err := createToken(*res, ctrl.jwtKey)
	if err != nil {
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
