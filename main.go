package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP Server on", getRunAddr())
	http.ListenAndServe(getRunAddr(), http.HandlerFunc(rootHandler))
}

func getRunAddr() string {
	return ":8000"
}
