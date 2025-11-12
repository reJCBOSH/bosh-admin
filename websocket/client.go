package websocket

import (
	"sync"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	writeWait      = 10 * time.Second    // 写入超时时间
	pongWait       = 60 * time.Second    // 心跳超时时间
	pingPeriod     = (pongWait * 9) / 10 // 心跳发送间隔时间
	maxMessageSize = 512 * 1024          //	最大消息长度
)

// Client  定义WebSocket客户端
type Client struct {
	ClientID        string             // 客户端ID
	hub             *Hub               // 所属Hub
	conn            *websocket.Conn    // WebSocket连接
	send            chan []byte        // 发送消息通道
	mutex           sync.Mutex         // 互斥锁
	active          bool               // 是否活跃
	connectAt       int64              // 连接时间戳
	lastHeartBeatAt int64              // 最后心跳时间戳
	Username        string             // 用户名
	logger          *zap.SugaredLogger // 日志
}

// NewClient 创建新客户端
func NewClient(hub *Hub, conn *websocket.Conn, username string) *Client {
	clientID, _ := random.UUIdV4()
	return &Client{
		ClientID:        clientID,
		hub:             hub,
		conn:            conn,
		send:            make(chan []byte, 256),
		active:          true,
		connectAt:       time.Now().Unix(),
		lastHeartBeatAt: time.Now().Unix(),
		Username:        username,
	}
}

// readPump  读取客户端消息
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.updateHeartBeat()
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.hub.logger.Errorf("error: %s", err.Error())
			}
			break
		}
		var msg Message
		err = msg.FromJson(message)
		if err != nil {
			c.hub.logger.Errorf("消息反序列化失败: %s", err.Error())
			break
		}
		if msg.Type == "ping" {
			respMsg := NewMessage(c.Username, "heartbeat", "heartbeat", "pong")
			resMessage, _ := respMsg.ToJson()
			c.send <- resMessage
		}
	}
}

// writePump  写入客户端消息
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			if err := w.Close(); err != nil {
				c.hub.logger.Errorf("error: %s", err.Error())
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) updateHeartBeat() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.lastHeartBeatAt = time.Now().Unix()
}

func (c *Client) isActive() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.active
}

func (c *Client) setActive(active bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.active = active
}
