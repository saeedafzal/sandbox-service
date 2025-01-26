package api

import (
	"net/http"

	"github.com/saeedafzal/sandbox-service/api/handlers"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	// Handlers
	healthHandler := handlers.HealthHandler{}

	// Routes
	mux.HandleFunc("GET /", healthHandler.GetVersion)

	return mux
}
