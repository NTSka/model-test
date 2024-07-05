package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"test-model/cmd/build"
	"test-model/cmd/step1/svc"
)

func main() {
	v := viper.New()
	v.SetConfigType("yaml")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	v.SetConfigFile(path.Join(wd, "local.yaml"))

	if err = v.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "viper.ReadInConfig"))
	}

	log.Println("Config read")

	producer, err := build.Producer(v)
	if err != nil {
		log.Fatal(errors.Wrap(err, "build.Producer"))
	}

	consumer, err := build.Consumer(v)
	if err != nil {
		log.Fatal(errors.Wrap(err, "build.Consumer"))
	}

	cfg := svc.NewConfig(v)
	service := svc.NewSvc(cfg, consumer, producer)

	log.Println("Starting service")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	err = service.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err = service.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("Stopping service")

	cancel()
	if err = producer.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("Producer stop")

	if err = consumer.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("Consumer stop")

	log.Println("Service finished")
}
