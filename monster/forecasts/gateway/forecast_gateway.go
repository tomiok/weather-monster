package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/monster/forecasts/models"
)

type ForecastGateway interface {
	GetForecastByCity(cityID int64) *models.ForecastResponse
}

type ForecastGtw struct {
	ForecastStorage
}

// GetForecastByCity given a cityID, get the avg for min and max temperatures in the las 24 hours. Also the number
// of samples will be returned.
func (g *ForecastGtw) GetForecastByCity(cityID int64) *models.ForecastResponse {
	return g.getForecastByCityID(cityID)
}

// NewForecastGateway easy way to instance a new gateway
func NewForecastGateway(client *database.SQLClient) ForecastGateway {
	return &ForecastGtw{&ForecastStg{client}}
}
