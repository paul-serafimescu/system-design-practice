package service

import (
	"context"
	"fmt"
	"http-server/database"
	"http-server/models"
	"http-server/repository"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func ValidateRegistrationRequest(r *models.ServiceRegistrationRequest) bool {
	return true // TODO: implement some extra checks to make sure it's all good
}

func GetChatWebsocketServer() (*models.RegisteredService, error) {
	servers, err := repository.GetAllAvailableServicesOfType(models.Websocket)
	if err != nil {
		return nil, err
	}

	return &servers[0], nil // change this to actually load balance lol
}

func FlagServiceAsDown(serviceId string) error {
	_, err := repository.UpdateServiceStatus(serviceId, models.Down)

	return err
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

	expiration := 300 * time.Second
	if err := cache.Expire(ctx, hkey, expiration).Err(); err != nil {
		return nil, err
	}

	log.Info().Msgf("websocket service with id: %s has been registered and will expire in %s", serviceId, expiration.String())

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

	log.Info().Msgf("successfully deleted")

	return nil
}

func RefreshServiceStatus(serviceId string) error {
	ctx := context.Background()
	cache := database.GetCache()

	service, err := repository.GetServiceByServiceId(serviceId)
	if err != nil {
		return err
	}

	if service.Status == models.Down {
		return fmt.Errorf("cannot refresh status: %s", "TTL expired")
	}

	switch service.Type {
	case models.Websocket:
		res, err := cache.Expire(ctx, fmt.Sprintf("websocket:%s", serviceId), 5*time.Minute).Result()

		if err != nil {
			return err
		}

		if !res {
			return fmt.Errorf("cannot refresh status: %s", "TTL expired")
		}
	default:
		return fmt.Errorf("unknown or invalid service type")
	}

	return nil
}
