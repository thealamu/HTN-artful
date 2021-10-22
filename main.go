package main

import (
	"log"
	"net/http"
)

func main() {
	if err := start(); err != nil {
		log.Println(err)
	}
}

func start() error {
	log.Println("Starting HTTP Server on", getRunAddr())
	mux := http.DefaultServeMux
	mux.HandleFunc("", rootHandler)
	return http.ListenAndServe(getRunAddr(), mux)
}

func getRunAddr() string {
	return ":8000"
}
