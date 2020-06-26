package config

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

// Configuration contains application configuration
type Configuration struct {
	Region         string
	HTTPPort       string        `mapstructure:"http_port" validate:"required"`
	AppName        string        `mapstructure:"app_name" validate:"required"`
	Env            string        `mapstructure:"environment" validate:"required"`
	LogLevel       string        `mapstructure:"log_level" validate:"required"`
	PokeAPIHost    string        `mapstructure:"poke_api_host" validate:"required"`
	PokeAPITimeout time.Duration `mapstructure:"poke_api_timeout" validate:"required"`
}

// BindAddress generates address with listening port
func (app *Configuration) BindAddress() string {
	return fmt.Sprintf("0.0.0.0:%s", app.HTTPPort)
}

// LoadConfiguration loads the remote config
func LoadConfiguration(publicConfigFile string) (*Configuration, error) {
	public := viper.New()
	public.SetConfigFile(publicConfigFile)
	if err := public.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Configuration
	err := public.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	if err = validate.Struct(cfg); err != nil {
		return nil, err
	}

	public.WatchConfig()
	public.OnConfigChange(func(e fsnotify.Event) {
		if err := public.Unmarshal(&cfg); err != nil {
			log.Print("failed to update public config after hot reload", "err", err)
		}
	})

	return &cfg, nil
}
