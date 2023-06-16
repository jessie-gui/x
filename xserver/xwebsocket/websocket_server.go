package xwebsocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jessie-gui/x/xlog"
)

// WebSocketClient 客户端连接对象。
type WebSocketClient struct {
	conn *websocket.Conn
}

// NewWebSocketClient 新建一个客户端对象。
func NewWebSocketClient(w http.ResponseWriter, r *http.Request) (*WebSocketClient, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	client := &WebSocketClient{
		conn: conn,
	}

	return client, nil
}

// ReadMessage 从连接中读取消息。
func (c *WebSocketClient) ReadMessage() ([]byte, error) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		xlog.Error("Error reading message:", err)
		return nil, err
	}

	return message, nil
}

// ReadMessageWithType 从连接中读取带有类型的消息。
func (c *WebSocketClient) ReadMessageWithType() (int, []byte, error) {
	ty, message, err := c.conn.ReadMessage()
	if err != nil {
		xlog.Error("Error reading messageWithType:", err)
		return 0, nil, err
	}

	return ty, message, nil
}

// WriteMessage 写入消息到连接中。
func (c *WebSocketClient) WriteMessage(message []byte) error {
	err := c.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		xlog.Error("Error writing message:", err)
		return err
	}

	return nil
}

// WriteMessageWithType 写入带有类型的消息到连接中。
func (c *WebSocketClient) WriteMessageWithType(msgType int, message []byte) error {
	err := c.conn.WriteMessage(msgType, message)
	if err != nil {
		xlog.Error("Error writing messageWithType:", err)
		return err
	}

	return nil
}

// Close 关闭连接。
func (c *WebSocketClient) Close() error {
	err := c.conn.Close()
	if err != nil {
		xlog.Error("Error closing connection:", err)
		return err
	}

	return nil
}
