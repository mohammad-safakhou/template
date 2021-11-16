package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func registerRoutes(e *echo.Echo, vConfig *viper.Viper, psqlDB  *sql.DB, logger  *zap.SugaredLogger) {
	controllers := NewRestControllers(vConfig, psqlDB, logger)
	v1 := e.Group("/v1")

	v1.GET("/hello", controllers.Hello)
}
