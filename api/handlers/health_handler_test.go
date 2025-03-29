package handlers_test

import (
	"net/http"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestHealthHandlerGetVersion200(t *testing.T) {
	s := setup()
	defer s.Close()

	res := GET(s, "/")
	defer res.Body.Close()

	assert.Equals(t, http.StatusOK, res.StatusCode)
}
