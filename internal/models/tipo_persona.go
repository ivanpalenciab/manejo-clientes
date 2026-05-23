package models

type TipoPersona struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Tipo string `gorm:"size:100;not null" json:"tipo_persona"`
}

func (TipoPersona) TableName() string {
	return "tipos_persona"
}
