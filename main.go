package main

import (
	"runtime"
	"strings"

	"api.inno/glo-profile-service/router"
	"api.inno/glo-profile-service/service"
	"api.inno/glo-profile-service/utility"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func init() {
	runtime.GOMAXPROCS(1)

}

func init() {
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.MergeConfig(strings.NewReader(viper.GetString("configs"))); err != nil {
		log.Panic(err.Error())
	} else {
		log.Info("loaded config " + viper.GetString("app.name"))
	}
	log.Info(viper.AllSettings())
}

func main() {
	r := router.InitRoute()
	service.InitRedisPoolService()
	utility.AESLoadKey()
	// utility.CallSOAPClientSteps()
	log.Info("started application listening on port " + viper.GetString("app.port"))
	r.Logger.Fatal(r.Start(":" + viper.GetString("app.port")))
}
