package models

type Ciudad struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `gorm:"size:100;not null" json:"nombre"`

	PaisID uint `gorm:"not null" json:"pais_id"`
	Pais   Pais `gorm:"foreignKey:PaisID"`
}
