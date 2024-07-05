package storage

import (
	"github.com/spf13/viper"
)

type Params struct {
	Dsn string
}

func NewConfig(v *viper.Viper) *Params {
	return &Params{
		Dsn: v.GetString("clickhouse.dsn"),
	}
}
