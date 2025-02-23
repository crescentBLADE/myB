package websocket

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	conn          *websocket.Conn
	url           string
	subscriptions map[string]func([]byte)
	mu            sync.RWMutex
}

func NewWsClient(streamURL string) *WsClient {
	return &WsClient{
		url:           streamURL,
		subscriptions: make(map[string]func([]byte)),
	}
}

func (c *WsClient) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	go c.readMessages()
	return nil
}

func (c *WsClient) Subscribe(stream string, handler func([]byte)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.subscriptions[stream] = handler
	c.conn.WriteJSON(map[string]interface{}{
		"method": "SUBSCRIBE",
		"params": []string{stream},
		"id":     time.Now().Unix(),
	})
}

func (c *WsClient) readMessages() {
	// for {
	//     _, msg, err := c.conn.ReadMessage()
	//     if err != nil {
	//         // 处理重连逻辑
	//         return
	//     }
	//     // 根据stream路由处理
	//     c.routeMessage(msg)
	// }
}
