package main

import (
	"chat/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", websocket.WsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
