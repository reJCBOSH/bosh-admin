package sse

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// SSEClient SSE客户端连接
type SSEClient struct {
	ClientID string
	Writer   gin.ResponseWriter
	Notify   <-chan bool
}

// SSEService 管理SSE客户端连接
type SSEService struct {
	clients map[string]*SSEClient
	mutex   sync.Mutex
}

// NewSSEService 创建新的SSEService实例
func NewSSEService() *SSEService {
	return &SSEService{
		clients: make(map[string]*SSEClient),
	}
}

// RegisterClient 注册一个新的SSE客户端
func (s *SSEService) RegisterClient(clientID string, writer gin.ResponseWriter, notify <-chan bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clients[clientID] = &SSEClient{
		ClientID: clientID,
		Writer:   writer,
		Notify:   notify,
	}
}

// UnregisterClient 注销一个SSE客户端
func (s *SSEService) UnregisterClient(clientID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clients, clientID)
}

// BroadcastEvent 广播事件给所有已注册的SSE客户端
func (s *SSEService) BroadcastEvent(eventType string, data string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, client := range s.clients {
		event := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, data)
		_, err := client.Writer.Write([]byte(event))
		if err != nil {
			// 如果写入失败，移除客户端
			s.UnregisterClient(client.ClientID)
			continue
		}
		client.Writer.Flush()
	}
}

// BroadcastEventTo 广播事件给指定的SSE客户端
func (s *SSEService) BroadcastEventTo(clientID string, eventType string, data string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if client, ok := s.clients[clientID]; ok {
		event := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, data)
		_, err := client.Writer.Write([]byte(event))
		if err != nil {
			s.UnregisterClient(client.ClientID)
			return
		}
		client.Writer.Flush()
	}
}

func (s *SSEService) HandleSSE(writer gin.ResponseWriter, clientID string) {
	// 设置SSE响应头
	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 创建通知通道，用于接收客户端的请求
	notify := writer.CloseNotify()
	// 注册SSE客户端
	s.RegisterClient(clientID, writer, notify)
	// 保持连接
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-notify:
			// 客户端请求关闭连接
			s.UnregisterClient(clientID)
			return
		case <-ticker.C:
			// 定时发送心跳包
			event := "event: ping\ndata: \n\n"
			_, err := writer.Write([]byte(event))
			if err != nil {
				s.UnregisterClient(clientID)
				return
			}
			writer.Flush()
		default:
			time.Sleep(time.Second)
		}
	}
}
