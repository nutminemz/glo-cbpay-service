package model

type PeriodDayRequest struct {
}

type PeriodDayResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Period `json:"data,omitempty"`
}

type Period struct {
	ID        string `json:"id,omitempty" redis:"id"`
	PeriodDt  string `json:"period_dtm,omitempty" redis:"period_dtm"`
	CreatedAt string `json:"creted_at,omitempty" redis:"creted_at"`
	UpdatedAt string `json:"updated_at,omitempty" redis:"updated_at"`
}
