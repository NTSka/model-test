package svc

import "github.com/spf13/viper"

type Config struct {
	XmlCount  int
	JSONCount int
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		XmlCount:  v.GetInt("svc.xml_count"),
		JSONCount: v.GetInt("svc.json_count"),
	}
}
