package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/monster/forecasts/gateway"
	"net/http"
	"strconv"
)

type ForecastHandler struct {
	gateway.ForecastGateway
}

func (h *ForecastHandler) GetForecastByCityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cityID, err := strconv.ParseInt(chi.URLParam(r, "cityID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}
	res := h.GetForecastByCity(cityID)

	if res == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "forecast unavailable for given city"})
	}
	json.NewEncoder(w).Encode(&res)

}

func NewForecastHandler(client *database.SQLClient) *ForecastHandler {
	return &ForecastHandler{gateway.NewForecastGateway(client)}
}
