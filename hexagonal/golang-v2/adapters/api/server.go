package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"template/config"
	"template/logger"
	"template/utils"
)

func StartRestServer() {
	l := logger.ZapLogger()
	zLogger := l.Sugar()
	vConfig, err := config.ViperConfig()
	if err != nil {
		panic(fmt.Sprintf("initializing viper failed %v", err))
	}

	psqlCon, err := utils.PostgresConnection(vConfig)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	registerMiddlewares(e, zLogger)
	registerRoutes(e, vConfig, psqlCon, zLogger)

	zLogger.Fatal(e.Start(":" + vConfig.GetString("server.port")))
}

func registerMiddlewares(e *echo.Echo, logger *zap.SugaredLogger) {

	// TODO change logger to custom logger
	e.Use(middleware.Logger())
	// TODO change recover to custom logger
	e.Use(middleware.Recover())
	t := loggerUtils{SugaredLogger: logger}
	e.HTTPErrorHandler = t.customZapHttpErrorHandler

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

type loggerUtils struct {
	*zap.SugaredLogger
}

func (z *loggerUtils) customZapHttpErrorHandler(err error, c echo.Context) {
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
