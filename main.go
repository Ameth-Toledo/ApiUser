package main

import (
	"UsersFree/src/core"
	"UsersFree/src/users/infrastructure/dependencies"
	"UsersFree/src/users/infrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func startServer() {
	for {
		log.Println("Iniciando servidor...")
		router := gin.Default()

		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		config.AllowHeaders = []string{"Authorization", "Content-Type"}
		config.ExposeHeaders = []string{"Content-Length", "Authorization"}
		config.MaxAge = 12 * time.Hour

		router.Use(cors.New(config))

		if err := initializeDependencies(router); err != nil {
			log.Fatalf("Error al inicializar dependencias: %v", err)
			return
		}
		go func() {
			if err := router.Run(":8080"); err != nil {
				log.Printf("Error en el servidor: %v", err)
			}
		}()
		time.Sleep(3 * time.Minute)
		log.Println("Reiniciando servidor...")
	}
}

func initializeDependencies(router *gin.Engine) error {

	pool := core.GetDBPool()

	createUserController, viewUserController, editUserController, deleteUserController, viewByIdUserController, loginController, userErr := dependencies.Init(pool)
	if userErr != nil {
		return userErr
	}

	routes.RegisterClientRoutes(router, createUserController, viewUserController, editUserController, deleteUserController, viewByIdUserController, loginController)

	return nil
}

func main() {
	startServer()
}
