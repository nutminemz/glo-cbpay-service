package service

import (
	"fmt"
	"math/big"
	"strings"

	"api.inno/glo-profile-service/model"
	"api.inno/glo-profile-service/utility"
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
)

func AddPayment(uuid string) model.AddPaymentResponse {
	redisKey := fmt.Sprintf(keyRedisProfile, uuid)
	re, err := redisPool.HGetRedis(redisKey)
	if err != nil {
		log.Infof("Error on init Redis %s", err.Error())
		response := model.AddPaymentResponse{
			Code:    1000,
			Message: err.Error(),
		}
		return response
	}
	res := model.Profile{}
	mapstructure.Decode(re, &res)
	var i big.Int
	i.SetString(strings.Replace(GenerateUUID(""), "-", "", 4), 16)

	txRef := i.String()[0:14]
	acc, _ := utility.Decrypt([]byte(res.CasaAc))
	cid, _ := utility.Decrypt([]byte(res.Cid))
	tel, _ := utility.Decrypt([]byte(res.MobileNo))

	code, info, err := utility.CallSOAPClientSteps(string(acc), txRef, string(cid), string(tel))
	qrs := model.AddPaymentResponse{}
	if err != nil {
		qrs.Code = model.StatusGenericError
		qrs.Message = err.Error()
	}
	if code == "0000" {
		qrs.Code = model.StatusSuccess
		qrs.Message = "success"
	} else {
		qrs.Code = model.StatusGenericError
		qrs.Message = info
	}
	return qrs
}
