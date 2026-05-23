package models

import "github.com/google/uuid"

type Usuario struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Nombre            string    `gorm:"size:100;not null" json:"nombre"`
	CorreoElectronico string    `gorm:"unique;not null" json:"correo_electronico"`
	Contrasena        string    `gorm:"not null" json:"password"`
	Key               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();uniqueIndex" json:"key"`
}
