package models

type Sucursal struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Nombre string `gorm:"size:100;not null" json:"nombre"`

	UbicacionID uint   `gorm:"not null" json:"ubicacion_id"`
	Ubicacion   Ciudad `gorm:"foreignKey:UbicacionID"`

	EmpresaID uint    `gorm:"not null" json:"empresa_id"`
	Empresa   Empresa `gorm:"foreignKey:EmpresaID"`
}

func (Sucursal) TableName() string {
	return "sucursales"
}
