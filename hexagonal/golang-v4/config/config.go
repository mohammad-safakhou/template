package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var AppConfig *config // global app config

type config struct {
	General   General   `mapstructure:"general"`   // general configs
	Databases Databases `mapstructure:"databases"` // databases configs
	Agents    Agents    `mapstructure:"agents"`
}

type General struct {
	Listen          string        `mapstructure:"listen"`           // rest listen port
	LogLevel        int8          `mapstructure:"log_level"`        // logger level
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"` // shutdown timeout
}

type Databases struct {
	Postgres Postgres `mapstructure:"postgres"` // postgres configs
}

type Postgres struct {
	Host         string        `mapstructure:"host"`           // postgres host
	Port         string        `mapstructure:"port"`           // postgres port
	User         string        `mapstructure:"user"`           // postgres user
	Pass         string        `mapstructure:"pass"`           // postgres pass
	DatabaseName string        `mapstructure:"database_name"`  // postgres database
	SslMode      string        `mapstructure:"ssl_mode"`       // postgres ssl mode
	MaxOpenConns int           `mapstructure:"max_open_conns"` // postgres max open connections
	MaxIdleConns int           `mapstructure:"max_idle_conns"` // postgres max idle connections
	Timeout      time.Duration `mapstructure:"timeout"`        // postgres timeout
}

type Agents struct {
}

// LoadConfig loads config from file
func LoadConfig(path string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name

	if path == "" {
		viper.AddConfigPath("./app/config")          // path to look for the config file in
		viper.AddConfigPath("./config")              // path to look for the config file in
		viper.AddConfigPath("$HOME/.config/daroghe") // call multiple times to add many search paths
		viper.AddConfigPath(".")                     // optionally look for config in the working directory
	} else {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	AppConfig = &config{}
	if err = viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
