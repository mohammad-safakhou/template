package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type RestControllers interface {
	Hello(ctx echo.Context) error
}

type restControllers struct {
	VConfig *viper.Viper
	PsqlDb  *sql.DB
	Logger  *zap.SugaredLogger
}

func NewRestControllers(vConfig *viper.Viper, psqlDB  *sql.DB, logger  *zap.SugaredLogger) RestControllers {
	return &restControllers{
		VConfig: vConfig,
		PsqlDb:  psqlDB,
		Logger:  logger,
	}
}
