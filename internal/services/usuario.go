package services

import (
	"my-api/internal/models"

	"my-api/internal/repository"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"

	"my-api/internal/requests"
)

func CrearUsuario(user *models.Usuario) *gorm.DB {

	hashedPassword, err := HashPassword(user.Contrasena)
	if err != nil {
		return &gorm.DB{Error: err}
	}

	user.Contrasena = hashedPassword

	return repository.CreateUser(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CambiarPassword(req requests.CambioPassword) *gorm.DB {
	hashedPassword, err := HashPassword(req.Password)

	if err != nil {
		return &gorm.DB{Error: err}
	}

	req.Password = hashedPassword

	return repository.CambiarPasswordRepository(req.Id, req.Password)
}

func UpdateUsuario(userId uint, user *requests.UpdateUsuarioRequest) *gorm.DB {

	hashedPassword, err := HashPassword(user.Contrasena)
	if err != nil {
		return &gorm.DB{Error: err}
	}

	user.Contrasena = hashedPassword

	return repository.UpdateUSer(userId, user)
}
