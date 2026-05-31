package handlers

import (
	"my-api/internal/models"

	//"my-api/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"

	"my-api/internal/repository"

	"strconv"
)

func CreateEmpresa(c *gin.Context) {
	var empresa models.Empresa

	if err := c.ShouldBindJSON(&empresa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := repository.CrearEmpresa(&empresa)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "empresa creada",
	})
}

func GetEmpresas(c *gin.Context) {
	var empresas []models.Empresa

	empresas, err := repository.GetEmpresas()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": empresas,
	})
}

func GetEmpresaByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id invalido",
		})
		return
	}
	empresa, err := repository.GetEmpresaByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": empresa,
	})
}
