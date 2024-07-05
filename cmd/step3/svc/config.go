package svc

import "github.com/spf13/viper"

type Config struct {
	Total     int
	ReportDir string
	Workers   int
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		Total:     v.GetInt("svc.total"),
		ReportDir: v.GetString("svc.report_dir"),
		Workers:   v.GetInt("svc.workers"),
	}
}
