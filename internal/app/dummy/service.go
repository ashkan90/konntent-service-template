package dummy

import (
	"context"
	"konntent-service-template/internal/app/dto/resource"
	"konntent-service-template/pkg/dummyclient"
	"konntent-service-template/pkg/event"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Handle(ctx context.Context, dto *event.DummyEvent) (resource.DummyResource, error)
}

type dummyService struct {
	logger *logrus.Logger
	client dummyclient.Client
}

func NewDummyService(l *logrus.Logger, c dummyclient.Client) Service {
	return &dummyService{
		logger: l,
		client: c,
	}
}

func (s *dummyService) Handle(ctx context.Context, dto *event.DummyEvent) (resource.DummyResource, error) {
	return resource.DummyResource{}, nil
}
