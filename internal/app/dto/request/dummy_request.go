package request

import (
	"konntent-service-template/pkg/event"
	"konntent-service-template/pkg/event/schema"
)

type DummyRequest struct {
	Data interface{} `json:"data" validate:"required"`
}

func (dr *DummyRequest) ToEvent() *event.DummyEvent {
	return &event.DummyEvent{EventData: dr.Data, EventType: schema.DummyEventType}
}
