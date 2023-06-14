package subscriber

import (
	"context"

	"github.com/jessie-gui/x/xevent/consumer"
	"github.com/jessie-gui/x/xevent/event"
)

// Subject /**
type Subject interface {
	AddConsumer(e event.EventType, consumer consumer.Consumer)
	RemoveConsumer(e event.EventType, consumer consumer.Consumer)
	NotifyConsumer(ctx context.Context, e event.Event)
}
