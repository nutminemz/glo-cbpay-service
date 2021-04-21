package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	handler "api.inno/glo-profile-service/handler"
)

func InitRoute() *echo.Echo {
	api := echo.New()

	api.HideBanner = true
	api.HidePort = true

	api.Pre(middleware.RemoveTrailingSlash())
	api.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Skipper: func(c echo.Context) bool {
				if c.Request().URL.Path == "/health" {
					return true
				}
				return false
			},
		},
	))

	if viper.GetString("log.level") == "debug" {
		api.Logger.SetLevel(log.DEBUG)
	} else {
		api.Logger.SetLevel(log.INFO)
	}

	api.Use(middleware.RequestID())
	api.Use(middleware.BodyLimit("5M"))
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
	}))

	api.Use(handler.ServerHeader)
	api.Use(handler.SetLogRequestID)

	api.GET("/health", handler.HealthCheck)

	api.GET("/onoff", handler.GetOnOffService)
	api.GET("/periodtime", handler.GetPeriodTime)
	api.GET("/memberstatus", handler.GetMemberStatus)
	api.GET("/periodday", handler.GetPeriodDay)
	api.GET("/soldout", handler.GetSoldOutFlag)

	api.POST("/addpayment", handler.AddPayment)

	return api
}
