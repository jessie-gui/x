package xtcp

import (
	"context"
	"crypto/tls"

	"github.com/jessie-gui/x/xlog"
	"github.com/jessie-gui/x/xnet/xtcp"
)

// ServerOption 定义一个 TCP 服务选项类型。
type ServerOption func(s *Server)

// Address 配置服务监听地址。
func Address(address string) ServerOption {
	return func(s *Server) { s.address = address }
}

// TLSConfig 配置 TLS。
func TLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) { s.tlsConfig = c }
}

// Handler 配置处理器。
func Handler(handler func(conn *xtcp.Conn)) ServerOption {
	return func(s *Server) { s.handler = handler }
}

// Server 定义 TCP 服务包装器。
type Server struct {
	*xtcp.Server

	address   string           // 服务器监听地址。
	handler   func(*xtcp.Conn) // 连接处理器。
	tlsConfig *tls.Config      // TLS 配置。
}

// NewServer 新建 TCP 服务器。
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		address: ":0",
		handler: func(conn *xtcp.Conn) {},
	}
	for _, opt := range opts {
		opt(srv)
	}
	if srv.tlsConfig != nil {
		srv.Server = xtcp.NewServerTLS(srv.address, srv.tlsConfig, srv.handler)
	} else {
		srv.Server = xtcp.NewServer(srv.address, srv.handler)
	}
	return srv
}

// Start 启动 TCP 服务器。
func (s *Server) Start(ctx context.Context) (err error) {
	xlog.Infof("[TCP] server listening on %s", s.GetListenedAddress())
	return s.Run(ctx)
}

// Stop 停止 TCP 服务器。
func (s *Server) Stop(ctx context.Context) error {
	xlog.Info("[TCP] server stopping")
	return s.Close(ctx)
}
