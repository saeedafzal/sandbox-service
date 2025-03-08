package handlers

import (
	"net/http"

	"github.com/saeedafzal/sandbox-service/model"
	"github.com/saeedafzal/sandbox-service/render"
)

type NicknameHandler struct {}

func (_ NicknameHandler) SetNickname(w http.ResponseWriter, r *http.Request) {
	// Get nickname from request and validate
	nicknameRequest := model.NicknameRequest{}
	if err := render.Bind(r, &nicknameRequest); err != nil {
		render.JSON(w, http.StatusBadRequest, render.M{"message": err.Error()})
		return
	}

	// Check if name already exists
	// Set name
	// Return OK
}
