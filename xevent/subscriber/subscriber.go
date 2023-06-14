package subscriber

import (
	"context"
	"sync"

	"github.com/jessie-gui/x/xevent/consumer"
	"github.com/jessie-gui/x/xevent/event"
	"github.com/jessie-gui/x/xlog"
)

// Subscriber 事件订阅对象。
type Subscriber struct {
	eventConsumers map[event.EventType]map[string]consumer.Consumer
	mu             sync.Mutex
}

func New() *Subscriber {
	return &Subscriber{
		eventConsumers: make(map[event.EventType]map[string]consumer.Consumer),
	}
}

// AddConsumer 添加观察者。
func (s *Subscriber) AddConsumer(e event.EventType, c consumer.Consumer) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.eventConsumers[e]; !ok {
		s.eventConsumers[e] = make(map[string]consumer.Consumer)
	}

	if _, ok := s.eventConsumers[e][c.Name()]; ok {
		xlog.Errorf("consumer:%s exists.", c.Name())
		return
	}

	s.eventConsumers[e][c.Name()] = c
	xlog.Debugf("add consumer:%s", c.Name())
}

// RemoveConsumer 移除观察者。
func (s *Subscriber) RemoveConsumer(e event.EventType, c consumer.Consumer) {
	if _, ok := s.eventConsumers[e]; !ok {
		s.eventConsumers[e] = make(map[string]consumer.Consumer)
		return
	}

	delete(s.eventConsumers[e], c.Name())
	xlog.Debugf("remove consumer:%s", c.Name())
}

// NotifyConsumer 通知观察者。
func (s *Subscriber) NotifyConsumer(ctx context.Context, e event.Event) {
	consumers, ok := s.eventConsumers[e.EventType()]
	if !ok {
		return
	}

	for _, c := range consumers {
		c.Exec(ctx, e)
	}
}
