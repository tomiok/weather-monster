package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/temperatures/models"
	"time"
)

type TemperatureStorage interface {
	saveTemperature(cmd *models.RegisterTemperatureCMD) (*models.Temperature, error)
}

type TemperatureStg struct {
	*database.SQLClient
}

func (s *TemperatureStg) saveTemperature(cmd *models.RegisterTemperatureCMD) (*models.Temperature, error) {
	tx, err := s.Begin()

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}
	ts := time.Now().Unix()

	res, err := tx.Exec(`insert into temperature (city_id, min, max, timestamp) values (?, ?, ?, ? )`,
		cmd.CityID, cmd.Min, cmd.Max, ts)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()
	return &models.Temperature{
		ID:        lastID,
		CityID:    cmd.CityID,
		Min:       cmd.Min,
		Max:       cmd.Max,
		Timestamp: ts,
	}, nil
}
