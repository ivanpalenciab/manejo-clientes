package routes

import (
	"my-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UsuarioRoutes(router *gin.Engine) {
	router.POST("/create-usuario", handlers.CreateUser)
	router.GET("/usuarios", handlers.GetUsers)
	router.GET("/usuario/:id", handlers.GetUserByID)
	router.POST("/usuario/login", handlers.Login)
	router.PUT("/usuario/cambiar-password", handlers.CambiarPassword)
	router.PUT("/usuario/update/:id", handlers.UpdateUser)
	/*router.DELETE("/usuario/:id", handlers.DeleteUser)*/
}
