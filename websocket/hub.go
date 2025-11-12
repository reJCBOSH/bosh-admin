package websocket

import (
	"net/http"
	"sync"

	"github.com/duke-git/lancet/v2/random"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Hub struct {
	clients    map[*Client]bool   // 客户端集合
	broadcast  chan []byte        // 广播通道
	register   chan *Client       // 注册通道
	unregister chan *Client       // 注销通道
	mutex      sync.RWMutex       // 读写锁
	logger     *zap.SugaredLogger // 日志
}

func NewHub(logger *zap.SugaredLogger) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte, 1000),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		logger:     logger,
	}
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.register:
			h.registerClient(client)
		case client := <-h.unregister:
			h.unregisterClient(client)
		case message := <-h.broadcast:
			h.broadcastMessage(message)
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	h.mutex.Lock()
	h.clients[client] = true
	h.mutex.Unlock()
	h.logger.Infof("客户端%s已连接", client.ClientID)
}

func (h *Hub) unregisterClient(client *Client) {
	if _, ok := h.clients[client]; ok {
		h.mutex.Lock()
		delete(h.clients, client)
		h.mutex.Unlock()
		close(client.send)
		h.logger.Infof("客户端%s已断开", client.ClientID)
	}
}

func (h *Hub) broadcastMessage(message []byte) {
	var msg Message
	err := msg.FromJson(message)
	if err != nil {
		h.logger.Errorf("消息反序列化失败: %s", err.Error())
		return
	}
	if msg.MessageID == "" {
		msg.MessageID, _ = random.UUIdV4()
	}
	// 如果指定了用户，只发送给该用户
	if msg.Username != "" {
		h.mutex.RLock()
		defer h.mutex.RUnlock()
		for client := range h.clients {
			if client.Username == msg.Username {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
		return
	}

	// 广播给所有客户端
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}

func (h *Hub) HandleConnection(w http.ResponseWriter, r *http.Request, username string) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Errorf("websocket升级失败: %s", err.Error())
		return
	}
	client := NewClient(h, conn, username)
	h.register <- client

	go client.writePump()
	go client.readPump()
}

func (h *Hub) Broadcast(msg *Message) {
	message, err := msg.ToJson()
	if err != nil {
		h.logger.Errorf("消息序列化失败: %s", err.Error())
		return
	}
	h.broadcastMessage(message)
}

func (h *Hub) SendToClient(clientID string, msg *Message) {
	message, err := msg.ToJson()
	if err != nil {
		h.logger.Errorf("消息序列化失败: %s", err.Error())
		return
	}
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	for client := range h.clients {
		if client.ClientID == clientID {
			client.send <- message
			break
		}
	}
}

func (h *Hub) SendToUser(username string, msg *Message) {
	message, err := msg.ToJson()
	if err != nil {
		h.logger.Errorf("消息序列化失败: %s", err.Error())
		return
	}
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	for client := range h.clients {
		if client.Username == username {
			client.send <- message
			break
		}
	}
}
