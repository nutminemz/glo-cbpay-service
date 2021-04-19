package model

type PinRequest struct {
	Uuid string `json:"uuid" gorm:"column:uuid"`
	Pin  string `json:"pin" gorm:"column:pin"`
}

type PinResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Profile `json:"data,omitempty"`
}
