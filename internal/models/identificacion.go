package models

type Identificacion struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `gorm:"size:100;not null" json:"nombre"`
	CODE   string `gorm:"size:10" json:"code"`
}

func (Identificacion) TableName() string {
	return "identificacion"
}
