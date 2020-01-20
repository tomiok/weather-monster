package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/forecasts/models"
	"time"
)

type ForecastStorage interface {
	getForecastByCityID(cityID int64) *models.ForecastResponse
}

type ForecastStg struct {
	*database.SQLClient
}

func (s *ForecastStg) getForecastByCityID(cityID int64) *models.ForecastResponse {
	tx, err := s.Begin()

	if err != nil {
		logs.Log().Error(err.Error())
		return nil
	}

	var res models.ForecastResponse
	now, before := calculate24HoursBefore()
	err = tx.QueryRow(`select count(*) as total, avg(min) as min, avg(max) as max from temperature 
	where city_id = ? and timestamp between ? and ?`, cityID, before, now).
		Scan(&res.Sample, &res.Min, &res.Max)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}
	_ = tx.Commit()

	return &res
}

func calculate24HoursBefore() (now int64, before24Hs int64) {
	now = time.Now().Unix()
	before24Hs = now - 86400
	return
}
