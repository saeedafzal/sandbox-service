package model

import (
	"net/http"
	"errors"
)

type NicknameRequest struct {
	Nickname string
}

func (n NicknameRequest) Bind(_ *http.Request) error {
	if len(n.Nickname) == 0 {
		return errors.New("error.nickname.required")
	}
	return nil
}
