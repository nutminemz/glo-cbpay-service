package service

import (
	"strconv"
	"strings"
	"time"

	"api.inno/glo-profile-service/model"
	"github.com/labstack/gommon/log"
)

func GetOnOffService() model.OnOffResponse {

	// waiting for TODO kafka inject to Redis
	err := redisPool.SetRedis(keyRedisServiceOff, "N")
	err = redisPool.SetRedis(keyRedisSoldOutFlag, "N")
	err = redisPool.SetRedis(keyRedisPeriodRound1, "2021-05-01")
	err = redisPool.SetRedis(keyRedisPeriodRound2, "2021-05-16")
	err = redisPool.SetRedis(keyRedisPeriodDay, "4,11,20,30")
	//
	currentTime := time.Now()
	y, _, day := currentTime.Date()

	serviceOff, err := redisPool.GetRedis(keyRedisServiceOff)
	soldOut, err := redisPool.GetRedis(keyRedisSoldOutFlag)
	periodSet, err := redisPool.GetRedis(keyRedisPeriodDay)
	roundKey := keyRedisPeriodRound1
	periodSetArr := strings.Split(periodSet, ",")
	period := periodSetArr[0] + "-" + periodSetArr[1]
	// TODO verify logic
	if day > 11 && day < 20 {
		roundKey = keyRedisPeriodRound2
		period = periodSetArr[2] + "-" + periodSetArr[3]
	}
	round, err := redisPool.GetRedis(roundKey)
	if err != nil {
		log.Infof("Error on get Redis %s", err.Error())
		response := model.OnOffResponse{
			Code:    model.StatusGenericError,
			Message: err.Error(),
		}
		return response
	}
	res := model.ServiceOff{}

	res.ServiceOff = false
	res.SoldOut = false
	layout := "2006-01-02"
	tRound, err := time.Parse(layout, round)
	res.Round = tRound.Format("02 January 2006")
	currentYear := strconv.Itoa(y)
	res.Period = period + " " + currentTime.Month().String() + " " + currentYear
	res.Type = "B"
	if serviceOff == "Y" {
		res.ServiceOff = true
	}
	if soldOut == "Y" {
		res.SoldOut = true
	}
	offFlag := model.OnOffResponse{
		Code:    model.StatusSuccess,
		Message: "success",
		Data:    res,
	}
	if err != nil {
		response := model.OnOffResponse{
			Code:    model.StatusGenericError,
			Message: err.Error(),
		}
		return response
	}
	return offFlag
}
