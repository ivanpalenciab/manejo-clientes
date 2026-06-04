package requests

type LoginRequest struct {
	NombreUsuario string `json:"nombre_usuario"`
	Password      string `json:"password"`
}

type CambioPassword struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
}

type UpdateUsuarioRequest struct {
	ID                uint   `gorm:"primaryKey"`
	NombreUsuario     string `gorm:"unique;not null" json:"nombre_usuario"`
	Nombre            string `gorm:"size:100;not null" json:"nombre"`
	CorreoElectronico string `gorm:"unique;not null" json:"correo_electronico"`
	Contrasena        string `gorm:"not null" json:"password"`
}
