package controllers

import (
	"UsersFree/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *application.AuthService
}

func NewAuthController(authService *application.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Execute(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	response, err := ac.authService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  response["token"],
		"userId": response["userId"],
		"name":   response["name"],
		"email":  response["email"],
	})
}
