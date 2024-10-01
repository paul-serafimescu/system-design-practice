package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // don't care about CSRF right now
}

func WsConnectionHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // nil because no cookies (for now at least)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v", r)
	fmt.Println("Client connected")
	conn.Close()
}
