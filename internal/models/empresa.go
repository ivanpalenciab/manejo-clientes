package models

type Empresa struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Nombre         string `gorm:"size:100;not null" json:"nombre"`
	Direccion      string `gorm:"size:100" json:"direccion"`
	Identificacion int    `gorm:"unique;not null" json:"identificacion"`
	CiudadID       uint   `gorm:"not null" json:"ciudad_id"`
	Ciudad         Ciudad `gorm:"foreignKey:CiudadID"`

	IdentificacionPersonaPaisID uint                      `gorm:"not null" json:"identificacion_persona_pais_id"`
	IdentificacionPersonaPais   IdentificacionPersonaPais `gorm:"foreignKey:IdentificacionPersonaPaisID"`

	Logo string `gorm:"size:100" json:"logo"`
}

func (Empresa) TableName() string {
	return "empresas"
}
