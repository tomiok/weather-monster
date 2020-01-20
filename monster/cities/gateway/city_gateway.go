package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/monster/cities/models"
)

type CityGateway interface {
	SaveCity(cmd *models.CreateCityCMD) (*models.City, error)
	DeleteCity(cityID int64) *models.City
	UpdateCity(cmd *models.UpdateCityCMD) *models.City
}

type CityGtw struct {
	CityStorage
}

// SaveCity stores a city in the DB
func (g *CityGtw) SaveCity(cmd *models.CreateCityCMD) (*models.City, error) {
	return g.saveCity(cmd)
}

// DeleteCity delete the city in the DB, with an ID
func (g *CityGtw) DeleteCity(cityID int64) *models.City {
	return g.deleteCity(cityID)
}

// UpdateCity updates the city in the DB
func (g *CityGtw) UpdateCity(cmd *models.UpdateCityCMD) *models.City {
	return g.updateCity(cmd)
}

// NewCityGatewayeasy way to instance a new gateway
func NewCityGateway(client *database.SQLClient) CityGateway {
	return &CityGtw{&CityStg{client}}
}
