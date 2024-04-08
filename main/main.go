package main

import (
	"net/http"
	"go_api_get_parameter/handler"
)

func main() {
	http.HandleFunc("/", handler.Router)
	http.ListenAndServe(":3000", nil)
}
