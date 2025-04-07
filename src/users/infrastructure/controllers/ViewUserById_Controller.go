package controllers

import (
	"UsersFree/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewUserByIdController struct {
	usecase application.ViewByIdUser
}

func NewViewUserByIdController(usecase application.ViewByIdUser) *ViewUserByIdController {
	return &ViewUserByIdController{usecase: usecase}
}

func (vc_c *ViewUserByIdController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	client, err := vc_c.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	client.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}
