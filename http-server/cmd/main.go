package main

import (
	"context"
	"fmt"
	"http-server/api"
	"http-server/config"
	"http-server/database"
	"http-server/service"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	cfg := config.GetConfig()

	pg, err := database.ConnectToDB(context.Background(), cfg)
	if err != nil {
		log.Fatal().Msgf("%v", err)
		os.Exit(1)
	}

	defer pg.Close()

	if err := pg.Ping(context.Background()); err != nil {
		log.Fatal().Msgf("%v", err)
		os.Exit(1)
	} else {
		log.Info().Msg("Connected to database...")
	}

	cache := database.ConnectToCache(context.Background(), cfg)

	defer cache.Close()

	if _, err := cache.Ping(context.Background()); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	} else {
		log.Info().Msg("Connected to cache (redis)...")
	}

	cache.OnWebsocketExpiration = func(serviceId string) error {
		err := service.FlagServiceAsDown(serviceId)

		if err != nil {
			return err
		}

		log.Error().Msgf("service %s is DOWN", serviceId)

		return nil
	}

	cache.OnError = func(err error) {
		log.Error().Msgf("%s", err.Error())
	}

	go cache.HandleKeyExpiration()

	apiServer := api.CreateApiServer()
	apiServer.StartListening()
}
