package service

import (
	"context"
	"fmt"
	"http-server/database"
	"http-server/models"
	"http-server/repository"
	"time"

	"github.com/google/uuid"
)

func ValidateRegistrationRequest(r *models.ServiceRegistrationRequest) bool {
	return true // TODO: implement some extra checks to make sure it's all good
}

// TODO: right now we generate the service id and dump it in redis, but it would probably be good to keep a backup on postgres
func RegisterWebsocket(r *models.ServiceRegistrationRequest) (*string, error) {
	ctx := context.Background()
	cache := database.GetCache()

	serviceId := uuid.New().String()

	if _, err := repository.CreateNewRegisteredWebsocketService(serviceId, r.Hostname, r.Port); err != nil {
		return nil, err
	}

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

func DeregisterService(serviceId string) error {
	ctx := context.Background()
	cache := database.GetCache()

	serviceType, err := repository.UpdateServiceStatus(serviceId, models.Down)

	if err != nil {
		return err
	}

	switch serviceType {
	case models.Websocket:
		err := cache.HDel(ctx, fmt.Sprintf("websocket:%s", serviceId), "hostname", "port").Err()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid service type")
	}

	fmt.Println("successfully deleted")

	return nil
}
