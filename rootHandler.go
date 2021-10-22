package main

import (
	"fmt"
	"net/http"
)

var rootResponseMsg = "Nothing to see here."

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootResponseMsg)
}
