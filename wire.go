//go:build wireinject
// +build wireinject

package konntent_service_template

import (
	"konntent-service-template/internal/app"
	"konntent-service-template/internal/app/dummy"
	"konntent-service-template/internal/app/handler"
	"konntent-service-template/internal/app/orchestration"
	"konntent-service-template/internal/listener/consumer"
	"konntent-service-template/pkg/claimer"
	"konntent-service-template/pkg/dummyclient"
	"konntent-service-template/pkg/nrclient"
	"konntent-service-template/pkg/rabbit"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var serviceProviders = wire.NewSet(
	dummy.NewDummyService,
	consumer.NewDummyConsumerService,
)

var orchestratorProviders = wire.NewSet(
	orchestration.NewDummyOrchestrator,
)

var handlerProviders = wire.NewSet(
	handler.NewDummyHandler,
)

var allProviders = wire.NewSet(
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
)

func InitAll(
	l *logrus.Logger,
	mc dummyclient.Client,
	mqp rabbit.Client,
	jwtInstance claimer.Claimer,
	nrInstance nrclient.NewRelicInstance,
) app.Router {
	wire.Build(allProviders, app.NewRoute)
	return nil
}
