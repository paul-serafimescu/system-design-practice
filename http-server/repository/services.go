package repository

import (
	"context"
	"http-server/database"
	"http-server/models"
)

func CreateNewRegisteredWebsocketService(serviceId string, hostname string, port int) (*models.RegisteredService, error) {
	registeredService := models.RegisteredService{
		ID:       serviceId,
		Hostname: hostname,
		Port:     port,
		Type:     models.Websocket,
		Status:   models.Up,
	}

	sql := "INSERT INTO registered_service (registered_service_id, hostname, port, status, type) VALUES ($1, $2, $3, $4, $5) RETURNING created_at, updated_at"
	err := database.Get().QueryRow(context.Background(),
		sql,
		serviceId,
		hostname,
		port,
		models.Up,
		models.Websocket).Scan(
		&registeredService.CreatedAt,
		&registeredService.LastUpdatedAt)

	return &registeredService, err
}

func UpdateServiceStatus(serviceId string, newStatus models.RegisteredServiceStatus) (models.ServiceType, error) {
	var serviceType models.ServiceType

	sql := "UPDATE registered_service SET status = $1 WHERE registered_service_id = $2 RETURNING type"
	err := database.Get().QueryRow(context.Background(), sql, newStatus, serviceId).Scan(&serviceType)

	return serviceType, err
}

func GetAllAvailableServicesOfType(serviceType models.ServiceType) ([]models.RegisteredService, error) {
	services := make([]models.RegisteredService, 0)
	sql := "SELECT registered_service_id, hostname, port, created_at, updated_at, status, type FROM registered_service WHERE status = UP AND type = $1"

	rows, err := database.Get().Query(context.Background(), sql, models.Websocket)
	if err != nil {
		return services, err
	}

	for rows.Next() {
		var service models.RegisteredService
		err := rows.Scan(&service.ID, &service.Hostname, &service.Port, &service.Status, &service.Type)

		if err != nil {
			return services, err
		}

		services = append(services, service)
	}

	return services, nil
}

func GetServiceByServiceId(serviceId string) (*models.RegisteredService, error) {
	var service models.RegisteredService
	sql := "SELECT registered_service_id, hostname, port, created_at, updated_at, status, type FROM registered_service WHERE registered_service_id = $1"

	err := database.Get().QueryRow(context.Background(), sql, serviceId).Scan(
		&service.ID,
		&service.Hostname,
		&service.Port,
		&service.CreatedAt,
		&service.LastUpdatedAt,
		&service.Status,
		&service.Type,
	)

	if err != nil {
		return nil, err
	}

	return &service, nil
}
