package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP Server on", getRunAddr())
	mux := http.DefaultServeMux
	mux.HandleFunc("", rootHandler)
	http.ListenAndServe(getRunAddr(), mux)
}

func getRunAddr() string {
	return ":8000"
}
