package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type ChatManager struct {
	connections map[string]*websocket.Conn
}

func NewChatManager() *ChatManager {
	return &ChatManager{
		connections: make(map[string]*websocket.Conn),
	}
}

func (c *ChatManager) Add(conn *websocket.Conn) string {
	username := generateUsername()
	c.connections[username] = conn
	fmt.Println("New user connected:", username)
	return username
}

func (c *ChatManager) Remove(username string) {
	delete(c.connections, username)
	fmt.Println("User disconnected:", username)
}

func (c *ChatManager) Broadcast(from string, msg []byte) {
	for username, conn := range c.connections {
		if username != from {
			conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func generateUsername() string {
	rand.Seed(time.Now().UnixNano())
	return "user" + fmt.Sprintf("%d", rand.Intn(100))
}
