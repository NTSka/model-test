package svc

import "github.com/spf13/viper"

type Config struct {
	Total int
	EPS   int
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		Total: v.GetInt("svc.total"),
		EPS:   v.GetInt("svc.eps"),
	}
}
