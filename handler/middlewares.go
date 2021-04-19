package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return next(c)
	}
}

func SetLogRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.SetHeader(`{"time":"${time_rfc3339_nano},"level":"${level}","file":"${short_file}:${line}"}`)
		return next(c)
	}
}
