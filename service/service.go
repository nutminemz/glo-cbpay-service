package service

import (
	"strings"

	"api.inno/glo-profile-service/utility"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

var redisPool *utility.RedisPool

var keyRedisProfile = "PL|%s"

var keyRedisServiceOff = "LOTTO_PERIOD_OFF"
var keyRedisPrice = "LOTTO_UNIT_PRICE"
var keyRedisPeriodDay = "LOTTO_PERIOD_SELL"
var keyRedisPeriodTime = "LOTTO_NORMAL_TIME_OPEN"

var keyRedisPeriodRound1 = "DTM_ROUND1"
var keyRedisPeriodRound2 = "DTM_ROUND2"

var keyRedisSoldOutFlag = "SOLD_OUT_FLAG"

func InitRedisPoolService() {
	log.Info("Initiating Redis Pool")
	redisPool = utility.InitPool(
		viper.GetString("redis.host"),
		viper.GetString("redis.port"),
		viper.GetString("redis.password"),
		viper.GetString("redis.database"),
		viper.GetInt("redis.poolMaxIdle"),
		viper.GetInt("redis.poolMaxActive"),
	)
}

func GenerateUUID(prefix string) string {
	uid := uuid.NewV4()
	return prefix + strings.ToUpper((strings.ReplaceAll(uid.String(), "-", "")))
}
