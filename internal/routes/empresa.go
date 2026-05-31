package routes

import (
	"my-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func EmpresaRoutes(router *gin.Engine) {
	router.POST("/create-empresa", handlers.CreateEmpresa)
	router.GET("/empresas", handlers.GetEmpresas)
	router.GET("/empresa/:id", handlers.GetEmpresaByID)
}
