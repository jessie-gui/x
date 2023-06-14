package events

import (
	"dy/event/event"
	"time"
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
