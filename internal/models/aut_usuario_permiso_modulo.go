package models

type AutPermisoModulo struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ModuloId uint   `gorm:"not null" json:"modulo_id"`
	Modulo   Modulo `gorm:"foreignKey:ModuloID"`

	UsuarioPermisoID  uint              `gorm:"not null" json:"usuario_permiso_id"`
	AutUsuarioPermiso AutUsuarioPermiso `gorm:"foreignKey:UsuarioPermisoID"`
}

func (AutPermisoModulo) TableName() string {
	return "aut_permisos_modulo"
}
