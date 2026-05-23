package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Identificacion string `gorm:"size:50;not null" json:"identificacion"`

	IdentificacionPersonaPaisId uint `gorm:"not null" json:"identificacion_persona_pais_id"`

	Nombre string `gorm:"size:100;not null" json:"nombre"`

	Apellido string `gorm:"size:100" json:"apellido"`

	Direccion string `gorm:"size:255" json:"direccion"`

	Telefono string `gorm:"size:20" json:"telefono"`

	Email string `gorm:"size:100" json:"email"`

	UsuarioCreateID uint `json:"usuario_create_id"`

	UsuarioUpdateID uint `json:"usuario_update_id"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Cliente) TableName() string {
	return "clientes"
}
