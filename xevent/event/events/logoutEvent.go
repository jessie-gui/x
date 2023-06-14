package events

import (
	"time"

	"github.com/jessie-gui/x/xevent/event"
)

// LogoutEvent /**
type LogoutEvent struct {
	UserID string
	Date   time.Time
}

func (l *LogoutEvent) EventType() event.EventType {
	return event.LogoutEventType
}

func (l *LogoutEvent) EventValue() event.Event {
	return l
}
