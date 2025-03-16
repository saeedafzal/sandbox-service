package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/saeedafzal/sandbox-service/api"
	"github.com/saeedafzal/sandbox-service/config"
	"github.com/saeedafzal/sandbox-service/store"
)

func main() {
	// Handle CLI flags
	if config.Flags() {
		return
	}

	// Start server
	mux := api.Init()
	port := store.GlobalStore.GetInt("port")
	server := http.Server{Addr: fmt.Sprintf(":%d", port), Handler: mux}

	go func() {
		slog.Info("Starting server:", "port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Error starting server:", "err", err)
			panic(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println()
	slog.Info("Performing graceful shutdown...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error during server shutdown:", "err", err)
		panic(err)
	}
	slog.Info("Server stopped.")
}
