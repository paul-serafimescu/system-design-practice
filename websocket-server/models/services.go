package models

type ServiceType int

const (
	Websocket ServiceType = iota
	Other
)

type ServiceRegistrationRequest struct {
	Hostname string      `json:"hostname"`
	Port     int         `json:"port"`
	Type     ServiceType `json:"type"`
	// TODO: anything else
}
