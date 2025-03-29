package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type WebSocketHandler struct {
	connections []*websocket.Conn
}

func NewWebSocketHandler() WebSocketHandler {
	return WebSocketHandler{make([]*websocket.Conn, 0)}
}

func (h WebSocketHandler) Connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info("Failed to upgrade websocket connection:", "err", err)
		return
	}
	defer c.Close()

	h.connections = append(h.connections, c)
	slog.Info("New websocket connection:", "count", len(h.connections))

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}

	slog.Info("WebSocket connection closing:", "count", len(h.connections))
}
