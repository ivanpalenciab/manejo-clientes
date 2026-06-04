package repository

import (
	"my-api/internal/database"

	"my-api/internal/models"

	"gorm.io/gorm"

	"context"

	"my-api/internal/requests"
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

func GetByUsername(username string) (models.Usuario, error) {
	usuario := models.Usuario{}
	connection := database.ConnectDB()

	result := connection.Where(&models.Usuario{NombreUsuario: username}).First(&usuario)

	if result.Error != nil {
		return usuario, result.Error
	}
	return usuario, nil
}

func CambiarPasswordRepository(id uint, password string) *gorm.DB {
	ctx := context.Background()

	connection := database.ConnectDB()
	//result := connection.Create(&usuario)
	result := connection.
		WithContext(ctx).
		Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("contrasena", password)

	return result

}

func UpdateUSer(id uint, usuario *requests.UpdateUsuarioRequest) *gorm.DB {
	updates := make(map[string]interface{})
	connection := database.ConnectDB()

	if usuario.Nombre != "" {
		updates["nombre"] = usuario.Nombre
	}

	if usuario.NombreUsuario != "" {
		updates["nombre_usuario"] = usuario.Nombre
	}

	if usuario.CorreoElectronico != "" {
		updates["correo_electronico"] = usuario.CorreoElectronico
	}

	if usuario.Contrasena != "" {
		updates["contrasena"] = usuario.Contrasena

	}

	return connection.
		Model(&models.Usuario{}).
		Where("id = ?", id).
		Updates(updates)
}
