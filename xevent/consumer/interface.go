package consumer

import (
	"context"
	"github.com/jessie-gui/x/xevent/event"
)

// Consumer 消费者接口
type Consumer interface {
	Exec(ctx context.Context, event event.Event) error
	Name() string
}
