package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/hyperjiang/gin-skeleton/model"
)

type AuthController struct {}

type SignupValidation struct {
	Email     string `json:"email" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required,min=6,max=20"`
}

func (ctrl *AuthController) DoLogin(c *gin.Context) {
	
}

func (ctrl *AuthController) DoSignup(c *gin.Context) {
	user := new(model.User)
	signupValidation := new(SignupValidation)

	if err := c.ShouldBindJSON(signupValidation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user.Email = signupValidation.Email
	user.Name = signupValidation.Name
	user.Password = signupValidation.Password

	if err := user.SignupJson(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	if err := jsonapi.MarshalPayload(c.Writer, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}