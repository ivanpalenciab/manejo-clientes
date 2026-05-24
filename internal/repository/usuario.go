package repository

import (
	"my-api/internal/database"
	"my-api/internal/models"

	"gorm.io/gorm"
)

func CreateUser(usuario *models.Usuario) *gorm.DB {
	connection := database.ConnectDB()
	result := connection.Create(&usuario)

	return result

}

func GetUsers() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	db := database.ConnectDB()

	result := db.Raw(`
		SELECT 
			id, 
			nombre_usuario,
			correo_electronico, 
			contrasena,
			"key",
			activo,
			created_at,
			updated_at
		FROM usuarios
	`).Scan(&usuarios)

	if result.Error != nil {
		return nil, result.Error
	}

	return usuarios, nil
}

func GetUser(userID int) (models.Usuario, error) {
	usuario := models.Usuario{}
	connection := database.ConnectDB()

	result := connection.First(&usuario, userID)

	if result.Error != nil {
		return usuario, result.Error
	}

	return usuario, nil
}
