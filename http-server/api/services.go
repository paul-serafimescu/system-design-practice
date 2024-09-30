package api

import (
	"encoding/json"
	"http-server/models"
	"http-server/service"
	"net/http"
)

func RegisterService(w http.ResponseWriter, r *http.Request) {
	var registrationRequest models.ServiceRegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&registrationRequest); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if !service.ValidateRegistrationRequest(&registrationRequest) {
		http.Error(w, "Forbidden Access", http.StatusForbidden)
		return
	}

	// choose which registration function we want to use based on service type
	var register func(r *models.ServiceRegistrationRequest) (*string, error)

	switch registrationRequest.Type {
	case models.Websocket:
		register = service.RegisterWebsocket
	default:
		http.Error(w, "Unknown Service Type", http.StatusBadRequest)
		return
	}

	serviceId, err := register(&registrationRequest)

	if err != nil {
		http.Error(w, "Failed to Register Service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(*serviceId))
}
