package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ApiServer struct {
	router *chi.Mux
}

func CreateApiServer() *ApiServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", CreateNewAccount)
		r.Post("/login", CreateSession)
	})

	r.Route("/services", func(r chi.Router) {
		r.Post("/register", RegisterService)
	})

	return &ApiServer{
		router: r,
	}
}

func (s *ApiServer) StartListening() {
	http.ListenAndServe(":8080", s.router)
}
