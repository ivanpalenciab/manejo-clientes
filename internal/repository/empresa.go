package repository

import (
	"my-api/internal/database"

	"my-api/internal/models"

	"gorm.io/gorm"
)

func CrearEmpresa(empresa *models.Empresa) *gorm.DB {
	connection := database.ConnectDB()
	result := connection.Create(&empresa)

	return result
}

func GetEmpresas() ([]models.Empresa, error) {
	var empresas []models.Empresa
	db := database.ConnectDB()

	result := db.Find(&empresas)

	if result.Error != nil {
		return nil, result.Error
	}

	return empresas, nil
}

func GetEmpresaByID(empresaID int) (models.Empresa, error) {
	empresa := models.Empresa{}
	connection := database.ConnectDB()

	result := connection.First(&empresa, empresaID)

	if result.Error != nil {
		return empresa, result.Error
	}

	return empresa, nil
}
