package models

type Cookie struct {
	Name     string `json:"name" validate:"required"`
	Value    string `json:"value" validate:"required"`
	Domain   string `json:"domain" validate:"required"`
	Path     string `json:"path" validate:"required"`
	MaxAge   string `json:"expires" validate:"required"`
	Secure   bool   `json:"secure" validate:"required"`
	HTTPOnly bool   `json:"httpOnly" validate:"required"`
}
