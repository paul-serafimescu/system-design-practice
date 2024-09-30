package main

import (
	"context"
	"fmt"
	"http-server/api"
	"http-server/config"
	"http-server/database"
	"os"
)

func main() {
	cfg := config.GetConfig()

	pg, err := database.ConnectToDB(context.Background(), cfg)
	if err != nil {
		// handle error in some way here
		os.Exit(1) // change this later
	}

	defer pg.Close()

	if err := pg.Ping(context.Background()); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1) // TODO
	} else {
		fmt.Println("Connected to database...")
	}

	cache := database.ConnectToCache(context.Background(), cfg)

	defer cache.Close()

	if _, err := cache.Ping(context.Background()); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Connected to cache (redis)...")
	}

	apiServer := api.CreateApiServer()
	apiServer.StartListening()
}
