package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP Server on", getRunAddr())

	// register handlers
	mux := http.DefaultServeMux
	mux.HandleFunc("/", rootHandler)

	chatManager := NewChatManager()
	mux.HandleFunc("/ws", socketHandler(chatManager))

	http.ListenAndServe(getRunAddr(), mux)
}

func getRunAddr() string {
	return ":8000"
}
