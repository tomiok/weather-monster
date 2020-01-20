package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/cities/gateway"
	"github.com/tomiok/weather-monster/monster/cities/models"
	"net/http"
	"strconv"
)

type CityHandler struct {
	gateway.CityGateway
}

func (h *CityHandler) CreateCityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd, err := parseCreateRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res, err := h.SaveCity(cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot save city"})
		return
	}

	json.NewEncoder(w).Encode(&res)

}

func (h *CityHandler) DeleteCityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cityID, err := strconv.ParseInt(chi.URLParam(r, "cityID"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.DeleteCity(cityID)

	if res == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot delete city"})
		return
	}

	json.NewEncoder(w).Encode(&res)
}

func (h *CityHandler) UpdateCityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cityID, err := strconv.ParseInt(chi.URLParam(r, "cityID"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	cmd, err := parseUpdateRequest(r, cityID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.UpdateCity(cmd)

	json.NewEncoder(w).Encode(&res)
}

func NewCityHandler(client *database.SQLClient) *CityHandler {
	return &CityHandler{gateway.NewCityGateway(client)}
}

func parseCreateRequest(r *http.Request) (*models.CreateCityCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.CreateCityCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}

	return &cmd, nil
}

func parseUpdateRequest(r *http.Request, id int64) (*models.UpdateCityCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd models.UpdateCityCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}
	cmd.ID = id
	return &cmd, nil
}
