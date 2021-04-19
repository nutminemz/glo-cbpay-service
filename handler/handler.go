package handler

import (
	"net/http"

	"api.inno/glo-profile-service/service"
	"api.inno/glo-profile-service/utility"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func HealthCheck(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, echo.Map{
		"message": "Service is Running !!",
	}, "	")
}

func GetOnOffService(c echo.Context) error {
	tokenAuth, err := utility.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	userId, err := utility.FetchAuth(tokenAuth)
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	log.Info(userId)
	// query := model.ProfileRequest{
	// 	Uuid: c.QueryParam("uuid"),
	// }
	resp := service.GetOnOffService()
	return c.JSONPretty(http.StatusOK, resp, "")
}

func GetPeriodTime(c echo.Context) error {
	tokenAuth, err := utility.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	userId, err := utility.FetchAuth(tokenAuth)
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	log.Info(userId)
	// query := model.ProfileRequest{
	// 	Uuid: c.QueryParam("uuid"),
	// }
	resp := service.GetPeriodTime()
	return c.JSONPretty(http.StatusOK, resp, "")
}

func GetMemberStatus(c echo.Context) error {
	tokenAuth, err := utility.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	userId, err := utility.FetchAuth(tokenAuth)
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	log.Info(userId)
	// query := model.ProfileRequest{
	// 	Uuid: c.QueryParam("uuid"),
	// }
	resp := service.GetMemberStatus()
	return c.JSONPretty(http.StatusOK, resp, "")
}

func GetPeriodDay(c echo.Context) error {
	tokenAuth, err := utility.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	userId, err := utility.FetchAuth(tokenAuth)
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	log.Info(userId)
	// query := model.ProfileRequest{
	// 	Uuid: c.QueryParam("uuid"),
	// }
	resp := service.GetPeriodDay()
	return c.JSONPretty(http.StatusOK, resp, "")
}

func GetSoldOutFlag(c echo.Context) error {
	tokenAuth, err := utility.ExtractTokenMetadata(c.Request())
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	userId, err := utility.FetchAuth(tokenAuth)
	if err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err, "")
	}
	log.Info(userId)
	// query := model.ProfileRequest{
	// 	Uuid: c.QueryParam("uuid"),
	// }
	resp := service.GetSoldOutFlag()
	return c.JSONPretty(http.StatusOK, resp, "")
}
