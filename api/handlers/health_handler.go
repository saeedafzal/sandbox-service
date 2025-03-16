package handlers

import (
	"net/http"

	"github.com/saeedafzal/sandbox-service/render"
	"github.com/saeedafzal/sandbox-service/store"
)

type HealthHandler struct{}

func (_ HealthHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, http.StatusOK, store.GlobalStore.GetString("version"))
}
