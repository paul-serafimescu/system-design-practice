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
		fmt.Println("error connecting")
		os.Exit(1) // TODO
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Println("Connected to database...")
	}

	apiServer := api.CreateApiServer()
	apiServer.StartListening()
}
