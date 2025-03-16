package handlers

import (
	"net/http"
	"log/slog"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type WebSocketHandler struct {}

func (_ WebSocketHandler) Connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info("Failed to upgrade websocket connection:", "err", err)
		return
	}
	defer c.Close()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}

	slog.Info("WebSocket connection closing.")
}
