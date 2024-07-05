package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"test-model/cmd/build"
	"test-model/cmd/producer/svc"
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

	producer, stopChan, err := build.Producer(v)
	if err != nil {
		log.Fatal(errors.Wrap(err, "build.Producer"))
	}

	cfg := svc.NewConfig(v)
	service := svc.NewSvc(cfg, producer)

	err = service.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	<-stopChan

	if err = producer.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("Success")
}
