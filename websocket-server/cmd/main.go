package main

import (
	"fmt"
	"websocket-server/config"
	"websocket-server/websocket"
)

func main() {
	cfg := config.GetConfig()

	wsServer := websocket.InitializeWebsocketServer()
	if err := wsServer.Start(cfg); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
