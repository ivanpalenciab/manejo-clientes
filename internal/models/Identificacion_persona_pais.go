package models

type IdentificacionPersonaPais struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PaisID uint `gorm:"not null" json:"pais_id"`
	Pais   Pais `gorm:"foreignKey:PaisID"`

	IdentificacionID uint           `gorm:"not null" json:"Identificacion_id"`
	Identificacion   Identificacion `gorm:"foreignKey:IdentificacionID"`

	TipoPersonaId uint        `gorm:"not null" json:"tipo_persona_id"`
	TipoPersona   TipoPersona `gorm:"foreignKey:TipoPersonaID"`
}

func (IdentificacionPersonaPais) TableName() string {
	return "identificacion_persona_pais"
}
