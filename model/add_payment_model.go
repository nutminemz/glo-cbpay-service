package model

type AddPaymentRequest struct {
}

type AddPaymentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
