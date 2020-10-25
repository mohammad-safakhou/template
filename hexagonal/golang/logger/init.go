package logger

import (
	"bizpooly/models/rest"
	"github.com/labstack/echo"
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

	if s, ok := err.(*rest.StandardHttpErrorResponse); ok {
		code = s.Code
		message = s.Message
	} else {
		z.Warn("try using frame works http error")
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(code, message)
	return
}
