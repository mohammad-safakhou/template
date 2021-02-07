package transport

import (
	"database/sql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ApplicationContext struct {
	VConfig *viper.Viper
	PsqlDb  *sql.DB
	Logger  *zap.SugaredLogger
}