package main

import (
	"doraserver/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/send", handlers.HandleSendMessage)
	http.HandleFunc("/messages", handlers.HandleGetMessages)

	fmt.Println("Server started. Port: 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
