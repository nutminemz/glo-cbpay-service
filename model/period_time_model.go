package model

type PeriodTimeRequest struct {
}

type PeriodTimeResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    Constants `json:"data,omitempty"`
}
