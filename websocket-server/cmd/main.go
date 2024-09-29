package main

import (
	"log"
	"net/http"
	"websocket-server/websocket"
)

func main() {
	http.HandleFunc("/ws", websocket.WsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
