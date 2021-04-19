package model

type SoldOutFlagRequest struct {
}

type SoldOutFlagResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Quota  `json:"data,omitempty"`
}

type Quota struct {
	ID         string `json:"id,omitempty" redis:"id"`
	PeriodDay  string `json:"period_day,omitempty" redis:"period_day"`
	SoldOutFlg string `json:"sold_out_flg,omitempty" redis:"sold_out_flg"`
	CreatedAt  string `json:"creted_at,omitempty" redis:"creted_at"`
	UpdatedAt  string `json:"updated_at,omitempty" redis:"updated_at"`
}
