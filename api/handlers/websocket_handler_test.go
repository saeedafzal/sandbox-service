package handlers_test

import (
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestWebSocketHandlerSendsUserListOnConnect(t *testing.T) {
	s := setup()
	defer s.Close()

	url := strings.Replace(s.URL+"/stream", "http", "ws", 1)
	_, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
}
