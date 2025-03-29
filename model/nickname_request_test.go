package model

import (
	"errors"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestNicknameRequestBindInvalidReturnsError(t *testing.T) {
	t.Parallel()

	req := NicknameRequest{"Name"}
	if err := req.Bind(nil); err != nil {
		assert.Equals(t, err.Error(), errors.New("error.nickname.required").Error())
	}
}

func TestNicknameRequestBindValid(t *testing.T) {
	t.Parallel()

	req := NicknameRequest{"Name"}
	err := req.Bind(nil)

	assert.NoError(t, err)
}
