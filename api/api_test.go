package api

import (
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestInit(t *testing.T) {
	mux := Init()
	assert.GreaterOrEqual(t, len(mux.Middlewares()), 1)
	assert.GreaterOrEqual(t, len(mux.Routes()), 1)
}
