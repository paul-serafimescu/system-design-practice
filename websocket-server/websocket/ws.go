package websocket

import (
	"fmt"
	"log"
	"net/http"
	"websocket-server/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // don't care about CSRF right now
}

type wsConnection struct {
	conn *websocket.Conn
}

func (c *wsConnection) sendMessage(msg *models.ChatMessage) {
	err := c.conn.WriteJSON(msg)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func WsConnectionHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil) // nil because no cookies (for now at least)

	if err != nil {
		log.Println(err)
		return
	}

	connection := wsConnection{c}

	// CONNECTION ESTABLISHED
	connection.sendMessage(&models.ChatMessage{
		Type: models.ClientHello,
		Payload: map[string]interface{}{
			"message": "hello",
		},
	})
}
