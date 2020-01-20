package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/webhook/gateway"
	"github.com/tomiok/weather-monster/monster/webhook/models"
	"net/http"
	"strconv"
)

type WebhookHandler struct {
	gateway.WebhookGateway
}

func (h *WebhookHandler) CreateWebhookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd, err := parseCreateForecastRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res, err := h.RegisterWebhook(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot register webhook"})
		return
	}

	json.NewEncoder(w).Encode(&res)
}

func (h *WebhookHandler) DeleteWebhookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	webhookID, err := strconv.ParseInt(chi.URLParam(r, "webhookID"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.DeleteWebhook(webhookID)

	json.NewEncoder(w).Encode(&res)
}

func NewWebhookHandler(client *database.SQLClient) *WebhookHandler {
	return &WebhookHandler{gateway.NewWebhookGateway(client)}
}

func parseCreateForecastRequest(r *http.Request) (*models.CreateWebhookCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.CreateWebhookCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}

	return &cmd, nil
}
