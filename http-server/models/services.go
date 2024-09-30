package models

import "time"

type ServiceType int

const (
	Websocket ServiceType = iota
	Other
)

// TODO: need some kind of "service secret" to verify that sender is actually a service
type ServiceRegistrationRequest struct {
	Hostname string      `json:"hostname"`
	Port     int         `json:"port"`
	Type     ServiceType `json:"type"`
	// TODO: anything else
}

type RegisteredServiceStatus string

const (
	Up   RegisteredServiceStatus = "UP"
	Down RegisteredServiceStatus = "DOWN"
	// TODO: add others
)

type RegisteredService struct {
	ID            string
	Hostname      string
	Port          int
	CreatedAt     time.Time
	LastUpdatedAt time.Time
	Status        RegisteredServiceStatus
	Type          ServiceType
}
