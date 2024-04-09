package main

import (
	"go_api_get_parameter/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/para", handler.Para)
	
	log.Println("Server is runnning on http://localhost:3000/para")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
