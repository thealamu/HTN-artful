package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketHandler(cm *ChatManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// upgrade the connection to websockets
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Error while upgrading connection to websockets", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// register this user with the chat manager
		username := cm.Add(conn)

		msgBuffer := make(chan []byte)
		done := readMessagesFromConn(conn, msgBuffer)

	outer:
		for {
			select {
			case msg := <-msgBuffer:
				// Broadcast to other connections
				cm.Broadcast(username, msg)
			case <-done:
				break outer
			}
		}
		cm.Remove(username)
	}
}

func readMessagesFromConn(conn *websocket.Conn, w chan []byte) (done chan bool) {
	done = make(chan bool)
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				done <- true
				break
			}
			w <- p
		}
	}()
	return
}
