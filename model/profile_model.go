package model

type ProfileRequest struct {
	Uuid string `json:"uuid" gorm:"column:uuid"`
}

type ProfileResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Profile `json:"data,omitempty"`
}

type Profile struct {
	ID        int    `json:"id,omitempty" redis:"id"`
	Uuid      string `json:"uuid,omitempty" redis:"uuid"`
	Cid       string `json:"cid,omitempty" redis:"cid"`
	Macid     string `json:"macid,omitempty" redis:"macid"`
	Firstname string `json:"firstname,omitempty" redis:"firstname"`
	Lastname  string `json:"lastname,omitempty" redis:"lastname"`
	MobileNo  string `json:"mobile_no,omitempty" redis:"mobile_no"`
	Email     string `json:"email,omitempty" redis:"email"`
	Pin       string `json:"pin,omitempty" redis:"pin"`
	CasaAc    string `json:"casa_ac,omitempty" redis:"casa_ac"`
	// GwalletAc       string `json:"gwallet_ac,omitempty" redis:"gwallet_ac"`
	Address         string `json:"address,omitempty" redis:"address"`
	Postcode        string `json:"postcode,omitempty" redis:"postcode"`
	DateOfBirth     string `json:"date_of_birth,omitempty" redis:"date_of_birth"`
	BiometricFlg    bool   `json:"biometric_flg,omitempty" redis:"biometric_flg"`
	NotificationFlg bool   `json:"notification_flg,omitempty" redis:"notification_flg"`
	AccessToken     string `json:"access_token,omitempty" redis:"access_token"`
}
