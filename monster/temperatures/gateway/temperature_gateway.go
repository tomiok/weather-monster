package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/monster/temperatures/models"
)

type TemperatureGateway interface {
	RegisterTemperature(cmd *models.RegisterTemperatureCMD) (*models.Temperature, error)
}

type TemperatureGtw struct {
	TemperatureStorage
}

// RegisterTemperature save the temperature with the current timestamp
func (g *TemperatureGtw) RegisterTemperature(cmd *models.RegisterTemperatureCMD) (*models.Temperature, error) {
	return g.saveTemperature(cmd)
}

// NewTemperatureGateway easy way to instance a new gateway
func NewTemperatureGateway(client *database.SQLClient) TemperatureGateway {
	return &TemperatureGtw{&TemperatureStg{client}}
}
