package model

type MemberStatusRequest struct {
}

type MemberStatusResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Member `json:"data,omitempty"`
}

type Member struct {
	ID          string `json:"id,omitempty" redis:"id"`
	Acct        string `json:"acct,omitempty" redis:"acct"`
	BlockFlg    string `json:"block_flg,omitempty" redis:"block_flg"`
	PeriodPaid  string `json:"period_paid,omitempty" redis:"period_paid"`
	PeriodDesc  string `json:"postoff_desc,omitempty" redis:"postoff_desc"`
	CitizenId   string `json:"citizen_id,omitempty" redis:"citizen_id"`
	RegisterFlg string `json:"register_flg,omitempty" redis:"register_flg"`
	CretedAt    string `json:"creted_at,omitempty" redis:"creted_at"`
	UpdatedAt   string `json:"updated_at,omitempty" redis:"updated_at"`
}
