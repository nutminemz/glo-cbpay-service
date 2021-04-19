package service

import (
	"fmt"

	"api.inno/glo-profile-service/model"
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
)

func GetPeriodDay() model.ProfileResponse {
	redisKey := fmt.Sprintf(keyRedisProfile, "xx")
	re, err := redisPool.HGetRedis(redisKey)
	if err != nil {
		log.Infof("Error on init Redis %s", err.Error())
		response := model.ProfileResponse{
			Code:    1000,
			Message: err.Error(),
		}
		return response
	}
	res := model.Profile{}
	mapstructure.Decode(re, &res)
	log.Print(res.Firstname)

	qrs := model.ProfileResponse{
		Code:    model.StatusSuccess,
		Message: "success",
		Data:    res,
	}
	return qrs
}
