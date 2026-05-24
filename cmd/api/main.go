package main

// @title API Gestion Negocios
// @version 1.0
// @description API para gestión de negocios por suscripción
// @host localhost:8080
// @BasePath /

import (
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "my-api/docs"

	"my-api/internal/database"
	"my-api/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env")
	}

	database.ConnectDB()

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UsuarioRoutes(r)

	log.Println("Servidor corriendo en puerto 8080")

	r.Run(":8080")
}
