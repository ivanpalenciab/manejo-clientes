package models

import "time"

type Suscripcion struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Costo       float64   `gorm:"type:decimal(10,2);not null"`
	FechaPago   time.Time `gorm:"not null"`
	FechaInicio time.Time `gorm:"not null"`
	FechaFin    time.Time `gorm:"not null"`
	Estado      bool      `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Suscripcion) TableName() string {
	return "suscripciones"
}
