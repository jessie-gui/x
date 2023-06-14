package event

// EventType 事件类型。
type EventType string

// Event 事件接口。
type Event interface {
	EventType() EventType
	EventValue() Event
}
