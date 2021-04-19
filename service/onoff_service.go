package service

import (
	"fmt"

	"api.inno/glo-profile-service/model"
	"github.com/labstack/gommon/log"
)

func GetOnOffService() model.OnOffResponse {
	redisKey := fmt.Sprintf(keyRedisOnOff)
	err := redisPool.SetRedis(redisKey, "Y")
	re, err := redisPool.GetRedis(redisKey)
	if err != nil {
		log.Infof("Error on init Redis %s", err.Error())
		response := model.OnOffResponse{
			Code:    1000,
			Message: err.Error(),
		}
		return response
	}
	res := model.ServiceOff{}
	if re == "Y" {
		res.ServiceOff = true
	} else {
		res.ServiceOff = false
	}

	offFlag := model.OnOffResponse{
		Code:    model.StatusSuccess,
		Message: "success",
		Data:    res,
	}
	return offFlag
}
