package subscriber

import (
	"context"
	"x/xevent/consumer"
	"x/xevent/event"
)

// Subject /**
type Subject interface {
	AddConsumer(e event.EventType, consumer consumer.Consumer)
	RemoveConsumer(e event.EventType, consumer consumer.Consumer)
	NotifyConsumer(ctx context.Context, e event.Event)
}
