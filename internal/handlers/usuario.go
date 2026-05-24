package handlers

import (
	//"log"
	"net/http"
	"time"

	"my-api/internal/repository"

	"github.com/gin-gonic/gin"

	"my-api/internal/models"

	"my-api/internal/services"

	"strconv"
)

// CreateUser godoc
// @Summary Crear usuario
// @Description Crea un nuevo usuario
// @Tags usuarios
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /create-usuario [post]
func CreateUser(c *gin.Context) {
	var user models.Usuario
	user.CreatedAt = time.Now()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := services.CrearUsuario(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "usuario creado",
	})
}

func GetUsers(c *gin.Context) {
	var usuarios []models.Usuario

	usuarios, err := repository.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usuarios,
	})
}

func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	usuario, err := repository.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usuario,
	})
}
