package main

import (
	"konntent-service-template/pkg/dummyclient"
	"konntent-service-template/pkg/rabbit"

	"github.com/sirupsen/logrus"
)

type server struct {
	logger      *logrus.Logger
	dummyClient dummyclient.Client
	mqInstance  rabbit.ConsumerInstance
}

func initLogger() *logrus.Logger {
	return logrus.New()
}
