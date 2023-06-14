package events

import (
	"time"

	"github.com/jessie-gui/x/xevent/event"
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
