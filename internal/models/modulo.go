package models

type Modulo struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `gorm:"size:100" json:"nombre"`
}
