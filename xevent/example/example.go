package example

import (
	"context"
	"time"
	"x/xevent/consumer"
	"x/xevent/event"
	"x/xevent/event/events"
	"x/xevent/subscriber"
)

/**
 *
 *
 * @author        Gavin Gui <guijiaxian@gmail.com>
 * @version       1.0.0
 * @copyright (c) 2022, Gavin Gui
 */
func main() {
	// 创建主题
	sub := subscriber.New()

	// 添加观察者
	user := NewUser()
	sub.AddConsumer(event.LoginEventType, consumer.ExecFunc(user.HandleLogin))
	sub.AddConsumer(event.LogoutEventType, consumer.ExecFunc(user.HandleLogout))

	// 创建用户事件对象
	loginEvent := &events.LoginEvent{UserID: "1", Date: time.Now()}
	logoutEvent := &events.LogoutEvent{UserID: "1", Date: time.Now()}

	// 事件回调(通知观察者)
	ctx := context.Background()
	sub.NotifyConsumer(ctx, loginEvent)
	sub.NotifyConsumer(ctx, logoutEvent)
}
