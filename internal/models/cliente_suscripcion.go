package models

type ClienteSuscripcion struct {
	Id        uint    `gorm:"primaryKey" json:"id"`
	ClienteID uint    `gorm:"not null" json:"cliente_id"`
	Cliente   Cliente `gorm:"foreignKey:ClienteId"`

	SucursalID uint
	Sucursal   Sucursal
}

func (ClienteSuscripcion) TableName() string {
	return "clientes_suscripciones"
}
