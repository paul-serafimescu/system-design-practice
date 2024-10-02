package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"websocket-server/config"
	"websocket-server/models"
)

// return true if we expired, false otherwise
func SendHeartbeat(serviceId string, hostname string, port string) bool {
	fmt.Printf("Sending heartbeat for service %s\n", serviceId)

	url := fmt.Sprintf("http://%s:%s/services/heartbeat", hostname, port)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Failed to send heartbeat: %v", err)
		return false
	}

	req.Header.Set("x-service-id", serviceId)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send heartbeat: %v", err)
		return false
	}

	defer resp.Body.Close()
	log.Printf("Heartbeat sent to %s, status code: %d", url, resp.StatusCode)

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
		fmt.Println(err.Error())
		return "", err
	}

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Successfully registered with service id: %s\n", string(respBody))

	defer resp.Body.Close()

	return string(respBody), nil
}
