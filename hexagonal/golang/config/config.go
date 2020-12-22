package config

import (
	"github.com/spf13/viper"
)

func ViperConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
