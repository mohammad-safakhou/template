package transport

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ApplicationContext struct {
	VConfig *viper.Viper
	PsqlDb  *sql.DB
	Logger  *zap.SugaredLogger
	Redis 	*redis.Client
}