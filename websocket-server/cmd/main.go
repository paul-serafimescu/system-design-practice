package main

import (
	"websocket-server/config"
	"websocket-server/websocket"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.GetConfig()

	wsServer := websocket.InitializeWebsocketServer()
	if err := wsServer.Start(cfg); err != nil {
		log.Error().Msgf("%s", err.Error())
	}
}
