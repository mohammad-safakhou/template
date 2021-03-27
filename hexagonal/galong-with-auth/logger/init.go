package logger

import (
	"backend-service/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func ZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}

type LoggerUtils struct {
	*zap.SugaredLogger
}

func (z LoggerUtils) CustomZapHttpErrorHandler(err error, c echo.Context) {
	var code = http.StatusInternalServerError
	var message interface{}

	if s, ok := err.(*utils.StandardHttpErrorResponse); ok {
		code = s.Code
		message = s.Message
	} else {
		z.Warn("try using http error frameworks")
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(code, message)
	return
}
