package example

import (
	"golang.org/x/net/context"
	"x/xevent/event"
	"x/xevent/event/events"
	"x/xlog"
)

/**
 *
 *
 * @author        Gavin Gui <guijiaxian@gmail.com>
 * @version       1.0.0
 * @copyright (c) 2022, Gavin Gui
 */
type user struct {
}

func NewUser() *user {
	return &user{}
}

// HandleLogin 监听登录消息。
func (u *user) HandleLogin(ctx context.Context, e event.Event) error {
	loginEvent, ok := e.EventValue().(*events.LoginEvent)
	if !ok {
		return nil
	}
	xlog.Infof("you login event at:%v\n", loginEvent.Date)

	return nil
}

// HandleLogout 监听退出消息。
func (u *user) HandleLogout(ctx context.Context, e event.Event) error {
	logoutEvent, ok := e.EventValue().(*events.LogoutEvent)
	if !ok {
		return nil
	}
	xlog.Infof("you logout event at:%v\n", logoutEvent.Date)

	return nil
}
