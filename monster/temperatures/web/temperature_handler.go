package web

import (
	"encoding/json"
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/temperatures/gateway"
	"github.com/tomiok/weather-monster/monster/temperatures/models"
	"net/http"
)

type TemperatureHandler struct {
	gateway.TemperatureGateway
}

func (h *TemperatureHandler) RegisterTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd, err := parseRegisterTemperatureCmd(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res, err := h.RegisterTemperature(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot register webhook"})
		return
	}

	json.NewEncoder(w).Encode(&res)
}

func NewTemperatureHandler(client *database.SQLClient) *TemperatureHandler {
	return &TemperatureHandler{gateway.NewTemperatureGateway(client)}
}

func parseRegisterTemperatureCmd(r *http.Request) (*models.RegisterTemperatureCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.RegisterTemperatureCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}

	return &cmd, nil
}
