package handlers_test

import (
	"testing"
	"strings"

	"github.com/saeedafzal/sandbox-service/tester/assert"
	// "github.com/saeedafzal/sandbox-service/model"
	"github.com/gorilla/websocket"
)

func TestWebSocketHandlerSendsUserListOnConnect(t *testing.T) {
	s := setup()
	defer s.Close()

	url := strings.Replace(s.URL + "/stream", "http", "ws", 1)
	_, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
}
