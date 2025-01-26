package api

import (
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestInit(t *testing.T) {
	mux := Init()
	assert.NotNil(t, mux)
}
