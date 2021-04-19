package model

type OnOffRequest struct {
}

type OnOffResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    ServiceOff `json:"data,omitempty"`
}

type Constants struct {
	ID          string `json:"id,omitempty" redis:"id"`
	ConfigKey   string `json:"config_key,omitempty" redis:"config_key"`
	ConfigValue string `json:"config_value,omitempty" redis:"config_value"`
	CreatedAt   string `json:"creted_at,omitempty" redis:"creted_at"`
	UpdatedAt   string `json:"updated_at,omitempty" redis:"updated_at"`
}

type ServiceOff struct {
	ServiceOff bool `json:"service_off,omitempty"`
}
