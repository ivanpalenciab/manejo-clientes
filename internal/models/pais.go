package models

type Pais struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `gorm:"size:100;not null" json:"nombre"`
	Codigo string `gorm:"size:3;not null" json:"codigo"`
}

func (Pais) TableName() string {
	return "paises"
}
