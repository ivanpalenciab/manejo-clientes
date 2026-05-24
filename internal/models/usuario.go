package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID                uint         `gorm:"primaryKey" json:"id"`
	NombreUsuario     string       `gorm:"unique;not null" json:"nombre_usuario"`
	Nombre            string       `gorm:"size:100;not null" json:"nombre"`
	CorreoElectronico string       `gorm:"unique;not null" json:"correo_electronico"`
	Contrasena        string       `gorm:"not null" json:"password"`
	Key               uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();uniqueIndex" json:"key"`
	ActivatedAt       sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt         time.Time    `gorm:"autoCreateTime"`
	UpdatedAt         time.Time    `gorm:"autoUpdateTime"`
	Activo            bool         `gorm:"default:false" json:"activo"`
}
