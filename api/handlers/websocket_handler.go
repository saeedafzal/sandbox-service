package handlers

import (
	"log/slog"
	"net/http"
	"context"
	"time"
	"crypto/rand"
	"encoding/hex"

	"github.com/saeedafzal/sandbox-service/store"
	"github.com/saeedafzal/sandbox-service/render"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type WebSocketHandler struct {
	connectionStore *store.Store[*websocket.Conn, struct{}]
	userStore *store.Store[string, struct{}]
}

func NewWebSocketHandler(userStore *store.Store[string, struct{}]) WebSocketHandler {
	return WebSocketHandler{
		connectionStore: store.New[*websocket.Conn, struct{}](),
		userStore: userStore,
	}
}

func (h WebSocketHandler) Connect(w http.ResponseWriter, r *http.Request) {
	// Accept new websocket connection
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{store.GetString("origin")},
	})
	if err != nil {
		slog.Warn("Failed to accept websocket connection:", "err", err)
		return
	}
	defer c.CloseNow()

	// Store connection
	h.connectionStore.Put(c, struct{}{})
	slog.Debug("New websocket connection:", "total", h.connectionStore.Length())

	// Send initial message (token + user list snapshot)
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	if err := wsjson.Write(ctx, c, render.M{
		"type": "initial",
		"token": h.generateRandomToken(),
		"users": h.userStore.Keys(),
	}); err != nil {
		slog.Error("Error sending user list to connection:", "err", err)
	}
	cancel()

	// Loop to keep connection open - ignore incoming messages
	ctx = context.Background()
	for {
		_, _, err := c.Reader(ctx)
		if err != nil {
			break
		}
	}

	h.connectionStore.Invalidate(c)
	slog.Debug("WebSocket connection closing:", "count", h.connectionStore.Length())
}

// Generates a random 64 digit token.
// TODO: Error handling
func (_ WebSocketHandler) generateRandomToken() string {
	b := make([]byte, 64)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
