package service

import (
	"context"
	"fmt"
	"http-server/database"
	"http-server/models"
	"time"

	"github.com/google/uuid"
)

func ValidateRegistrationRequest(r *models.ServiceRegistrationRequest) bool {
	return true // TODO: implement some extra checks to make sure it's all good
}

func RegisterWebsocket(r *models.ServiceRegistrationRequest) (*string, error) {
	ctx := context.Background()
	cache := database.GetCache()

	serviceId := uuid.New().String()

	hkey := fmt.Sprintf("websocket:%s", serviceId)
	if err := cache.HSet(ctx, hkey, map[string]interface{}{
		"hostname": r.Hostname,
		"port":     r.Port,
	}).Err(); err != nil {
		return nil, err
	}

	expiration := 5 * time.Minute
	if err := cache.Expire(ctx, hkey, expiration).Err(); err != nil {
		return nil, err
	}

	// TODO: change this message to a proper log
	fmt.Printf("websocket service with id: %s has been registered and will expire in 300 seconds\n", serviceId)

	return &serviceId, nil
}
