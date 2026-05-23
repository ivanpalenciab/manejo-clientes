package models

type ClienteSucursal struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ClienteID uint    `gorm:"not null" json:"cliente_id"`
	Cliente   Cliente `gorm:"foreignKey:ClienteId"`

	SucursalID uint     `gorm:"not null" json:"sucursal_id"`
	Sucursal   Sucursal `gorms:"foreignKey:SucursalID"`
}

func (ClienteSucursal) TableName() string {
	return "clientes_sucursales"
}
