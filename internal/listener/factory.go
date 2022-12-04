package listener

import (
	"konntent-service-template/internal/listener/consumer"
	"konntent-service-template/internal/listener/handler"
	"konntent-service-template/pkg/event"
	"konntent-service-template/pkg/event/schema"
	"konntent-service-template/pkg/eventmanager"
)

type eventHandlerFactory struct {
	service consumer.Service
}

func NewEventHandlerFactory(cs consumer.Service) eventmanager.IEventHandlerFactory {
	return &eventHandlerFactory{
		service: cs,
	}
}

func (eh *eventHandlerFactory) Make(e event.Factory) (eventmanager.EventHandler, error) {
	if e.Type() == schema.DummyEventType {
		return handler.NewDummyConsumerHandler(eh.service, e.Data().(*event.DummyEvent)), nil
	}

	return nil, schema.UnexpectedHandlerType
}
