package consumer

import (
	"context"
	"fmt"

	"github.com/jessie-gui/x/xevent/event"
)

// ExecFunc 消费者类型。
type ExecFunc func(ctx context.Context, event event.Event) error

// Exec 事件执行方法。
func (f ExecFunc) Exec(ctx context.Context, event event.Event) error {
	return f(ctx, event)
}

// Name 事件方法名。
func (f ExecFunc) Name() string {
	return fmt.Sprintf("%d", &f)
}
