package routes

import (
	"UsersFree/src/core/security"
	"UsersFree/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterClientRoutes(r *gin.Engine,
	createClientController *controllers.CreateUserController,
	viewClientController *controllers.ViewUserController,
	editClientController *controllers.EditUserController,
	deleteClientController *controllers.DeleteUserController,
	viewByIdClientController *controllers.ViewUserByIdController,
	loginController *controllers.AuthController,
) {
	r.POST("/clients", createClientController.Execute)
	r.GET("/clients", security.JWTMiddleware(), viewClientController.Execute)
	r.POST("/login", loginController.Execute)
	r.GET("/clients/:id", security.JWTMiddleware(), viewByIdClientController.Execute)
	r.PUT("/clients/:id", security.JWTMiddleware(), editClientController.Execute)
	r.DELETE("/clients/:id", security.JWTMiddleware(), deleteClientController.Execute)
}
