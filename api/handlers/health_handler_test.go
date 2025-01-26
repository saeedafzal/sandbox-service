package handlers_test

import (
	"testing"
	"net/http"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestHealthHandlerGetVersion200(t *testing.T) {
	s := setup()
	defer s.Close()

	res := GET(s, "/")
	defer res.Body.Close()

	assert.Equals(t, http.StatusOK, res.StatusCode)
}
