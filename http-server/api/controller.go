package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ApiServer struct {
	router *chi.Mux
}

func CreateApiServer() *ApiServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	corsOptions := cors.Options{
		// '*' allows all origins
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	r.Use(cors.Handler(corsOptions))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", CreateNewAccount)
		r.Post("/login", CreateSession)
	})

	r.Route("/services", func(r chi.Router) {
		r.Post("/register", RegisterService)
		r.Delete("/deregister/{serviceID}", DeregisterService)
		r.Get("/heartbeat", ReceiveHeartbeat)
	})

	return &ApiServer{
		router: r,
	}
}

func (s *ApiServer) StartListening() {
	http.ListenAndServe(":8080", s.router)
}
