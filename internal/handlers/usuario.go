package handlers

import (
	//"log"
	"net/http"
	"time"

	"my-api/internal/repository"

	"github.com/gin-gonic/gin"

	"my-api/internal/models"

	"my-api/internal/services"

	"my-api/internal/requests"

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

func Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request inválido",
		})
		return
	}

	usuario, err := repository.GetByUsername(req.NombreUsuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !services.CheckPassword(usuario.Contrasena, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "usuario o contraseña incorrectos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "credenciales válidas",
	})
}

func CambiarPassword(c *gin.Context) {
	var cambio requests.CambioPassword

	//user.UpdatedAt   = time.Now()
	if err := c.ShouldBindJSON(&cambio); err != nil {
		println("falló aquí")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := services.CambiarPassword(cambio)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error de servidor": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contraseña actualizada",
	})

}

func UpdateUser(c *gin.Context) {
	var user requests.UpdateUsuarioRequest

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	id := uint(id64)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := services.UpdateUsuario(id, &user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "usuario actualizado",
	})

}
