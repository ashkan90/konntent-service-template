package handler

import (
	"context"
	"konntent-service-template/internal/listener/consumer"
	"konntent-service-template/pkg/dummyclient/model"
	"konntent-service-template/pkg/event"
	"konntent-service-template/pkg/eventmanager"
	"konntent-service-template/pkg/utils"
)

type dummyConsumerHandler struct {
	service consumer.Service
	event   *event.DummyEvent
}

func NewDummyConsumerHandler(service consumer.Service, event *event.DummyEvent) eventmanager.EventHandler {
	return &dummyConsumerHandler{
		service: service,
		event:   event,
	}
}

func (o *dummyConsumerHandler) Handle(ctx context.Context) error {
	ctx = o.prepareHeaderMap(ctx)
	return o.service.DummyEvent(ctx, o.preparePayload())
}

func (o *dummyConsumerHandler) preparePayload() model.DummyRequest {
	return model.DummyRequest{
		Data: o.event.EventData,
	}
}

func (o *dummyConsumerHandler) prepareHeaderMap(ctx context.Context) context.Context {
	return context.WithValue(ctx, utils.HeaderMapCtx, o.event.EventHeaders)
}
