package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	addr := getRunAddr()
	fmt.Println("Starting HTTP Server on", addr)
	http.ListenAndServe(addr, http.HandlerFunc(rootHandler))
}

func getRunAddr() string {
	return ":8000"
}
