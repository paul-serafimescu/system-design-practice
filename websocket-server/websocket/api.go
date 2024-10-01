package websocket

import (
	"fmt"
	"net"
	"net/http"
	"websocket-server/config"
	"websocket-server/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Route("/ws", func(r chi.Router) {
		r.HandleFunc("/connect", WsConnectionHandler)
	})

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

func getPortListenerInRange(start, end int) (net.Listener, error) {
	for port := start; port <= end; port++ {
		listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
		if err == nil {
			return listener, nil
		}
	}
	return nil, fmt.Errorf("no available port found in range %d-%d", start, end)
}

func (wss *WebsocketServer) Start(cfg *config.Config) error {
	listener, err := getPortListenerInRange(8000, 8030) // for now
	if err != nil {
		panic(err)
	}

	tcpAddr := listener.Addr().(*net.TCPAddr)
	wss.hostname = tcpAddr.IP.String() // small issue: this returns :: but apparently docker-compose handles its own DNS? weird. so we only need port i guess
	wss.port = tcpAddr.Port

	service.RegisterService(cfg, wss.GetHostname(), wss.GetPort())

	fmt.Printf("Listening on port: %d\n", wss.port)

	return http.Serve(listener, nil)
}
