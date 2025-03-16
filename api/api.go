package api

import (
	"github.com/saeedafzal/sandbox-service/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	mux := chi.NewRouter()

	// Global middleware
	mux.Use(
		middleware.RequestID,
		middleware.Recoverer,
	)

	// Handlers
	healthHandler := handlers.HealthHandler{}
	nicknameHandler := handlers.NicknameHandler{}
	websocketHandler := handlers.WebSocketHandler{}

	// Routes
	mux.Get("/", healthHandler.GetVersion)
	mux.Post("/nickname", nicknameHandler.SetNickname)
	mux.Get("/stream", websocketHandler.Connect)

	return mux
}
