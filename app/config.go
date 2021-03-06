package app

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/go-ozzo/ozzo-validation"
)

var Config appConfig

type appConfig struct {
	ErrorFile  string `mapstructure:"error_file"`
	ServerPort string `mapstructure:"server_port"`
	DSN        string `mapstructure:"dsn"`
	Release    string `mapstructure:"release"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
		validation.Field(&config.Release, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("server_port", "")
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
