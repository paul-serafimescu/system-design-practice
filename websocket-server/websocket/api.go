package websocket

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"websocket-server/config"
	"websocket-server/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type WebsocketServer struct {
	router   *chi.Mux
	hostname string
	port     int
}

func InitializeWebsocketServer() *WebsocketServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/connect", WsConnectionHandler)

	return &WebsocketServer{
		router: r,
	}
}

func (wss *WebsocketServer) GetHostname() string {
	return wss.hostname
}

func (wss *WebsocketServer) GetPort() int {
	return wss.port
}

func getPublicHostname() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// Check if the address is an IP address and not a loopback
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // Ensure it's an IPv4 address
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no valid public IP address found")
}

func (wss *WebsocketServer) Start(cfg *config.Config) error {
	listener, err := net.Listen("tcp", ":9000") // for now
	if err != nil {
		log.Fatal().Msgf("%v", err)
		panic(err)
	}

	wss.hostname, _ = getPublicHostname()
	wss.port = listener.Addr().(*net.TCPAddr).Port

	serviceId, err := service.RegisterService(cfg, wss.GetHostname(), wss.GetPort())
	if err != nil {
		log.Fatal().Msgf("%v", err)
		panic(err)
	}

	ticker := time.NewTicker(120 * time.Second)
	go func() {
		for range ticker.C {
			isExpired := service.SendHeartbeat(serviceId, cfg.RegistryHost, cfg.RegistryPort)

			if isExpired {
				serviceId, err = service.RegisterService(cfg, wss.GetHostname(), wss.GetPort())

				if err != nil {
					log.Fatal().Msgf("%v", err)
					panic(err)
				}
			}
		}
	}()

	log.Info().Msgf("Listening on port: %d", wss.port)

	return http.Serve(listener, wss.router)
}
