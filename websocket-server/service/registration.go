package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"websocket-server/config"
	"websocket-server/models"

	"github.com/rs/zerolog/log"
)

// return true if we expired, false otherwise
func SendHeartbeat(serviceId string, hostname string, port string) bool {
	fmt.Printf("Sending heartbeat for service %s\n", serviceId)

	url := fmt.Sprintf("http://%s:%s/services/heartbeat", hostname, port)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Msgf("Failed to send heartbeat: %v", err)
		return false
	}

	req.Header.Set("x-service-id", serviceId)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Failed to send heartbeat: %v", err)
		return false
	}

	defer resp.Body.Close()
	log.Info().Msgf("Heartbeat sent to %s, status code: %d", url, resp.StatusCode)

	return resp.StatusCode >= 400
}

func RegisterService(cfg *config.Config, hostname string, port int) (string, error) {
	registryEndpoint := fmt.Sprintf("http://%s:%s/services/register", cfg.RegistryHost, cfg.RegistryPort)

	body, _ := json.Marshal(models.ServiceRegistrationRequest{
		Hostname: hostname,
		Port:     port,
		Type:     models.Websocket,
	})

	resp, err := http.Post(registryEndpoint, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	respBody, _ := io.ReadAll(resp.Body)
	log.Info().Msgf("Successfully registered with service id: %s", string(respBody))

	defer resp.Body.Close()

	return string(respBody), nil
}
