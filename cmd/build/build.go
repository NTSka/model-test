package build

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"test-model/pkg/amqp"
	"test-model/pkg/storage"
)

func Producer(v *viper.Viper) (amqp.Producer, error) {
	cfg := amqp.NewConfig(v)

	producer := amqp.NewProducer(cfg)

	err := producer.Init()
	if err != nil {
		return nil, errors.Wrap(err, "producer.Init")
	}

	return producer, err
}

func Consumer(v *viper.Viper) (amqp.Consumer, error) {
	cfg := amqp.NewConfig(v)

	producer := amqp.NewConsumer(cfg)

	err := producer.Init()
	if err != nil {
		return nil, errors.Wrap(err, "producer.Init")
	}

	return producer, err
}

func Clickhouse(v *viper.Viper) (storage.Clickhouse, error) {
	cfg := storage.NewConfig(v)

	db, err := storage.NewClickhouse(cfg)
	return db, errors.Wrap(err, "storage.NewClickhouse")
}
