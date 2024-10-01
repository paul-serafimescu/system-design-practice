package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"websocket-server/config"
	"websocket-server/models"
)

func RegisterService(cfg *config.Config, hostname string, port int) error {
	registryEndpoint := fmt.Sprintf("http://%s:%s/services/register", cfg.RegistryHost, cfg.RegistryPort)

	body, _ := json.Marshal(models.ServiceRegistrationRequest{
		Hostname: hostname,
		Port:     port,
		Type:     models.Websocket,
	})

	resp, err := http.Post(registryEndpoint, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Received from server: %s\n", string(respBody))

	defer resp.Body.Close()

	return nil
}
