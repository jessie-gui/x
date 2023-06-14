package events

import (
	"dy/event/event"
	"time"
)

// LoginEvent 定义登录事件详情
type LoginEvent struct {
	UserID string
	Date   time.Time
}

func (l *LoginEvent) EventType() event.EventType {
	return event.LoginEventType
}

func (l *LoginEvent) EventValue() event.Event {
	return l
}
