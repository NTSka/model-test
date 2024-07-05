package amqp

import (
	"github.com/spf13/viper"
)

type Config struct {
	DSN string
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		DSN: v.GetString("amqp.dsn"),
	}
}
