package config

import "C"
import (
	"fmt"
	"github.com/spf13/viper"
)

func PostgresURI(c *viper.Viper) string {
	host := c.GetString("database.psql.host")
	port := c.GetString("database.psql.port")
	user := c.GetString("database.psql.user")
	pass := c.GetString("database.psql.password")
	database := c.GetString("database.psql.name")
	sslmode := c.GetString("database.psql.sslmode")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, database, sslmode)
}
