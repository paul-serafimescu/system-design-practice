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

	if pg.Ping(context.Background()) != nil {
		fmt.Println("error connecting")
		os.Exit(1) // TODO
	} else {
		fmt.Println("connected!!")
	}

	apiServer := api.CreateApiServer()
	apiServer.StartListening()
}
