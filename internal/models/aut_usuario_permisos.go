package models

type AutUsuarioPermiso struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UsiarioID uint    `gorm:"not null" json:"usuario_id"`
	Usuario   Usuario `gorm:"foreignKey:UsusuarioID"`

	SucursalID uint     `gorm:"not null" json:"sucursal_id"`
	Sucursal   Sucursal `gorm:"foreignKey:SucursalID"`
}

func (AutUsuarioPermiso) TableName() string {
	return "aut_usuarios_permisos"
}
