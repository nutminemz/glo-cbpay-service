package service

import (
	"strings"

	"api.inno/glo-profile-service/utility"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

var redisPool *utility.RedisPool

var keyRedisOnOff = "ONOFF"
var keyRedisProfile = "PL|%s"
var keyRedisPin = "PI|%s"
var keyRedisAppVersion = "APPVERSION"

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
