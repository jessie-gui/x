package xserver

import (
	"context"
)

// Server 服务接口
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
