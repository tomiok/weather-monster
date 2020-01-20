package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/cities/models"
)

type CityStorage interface {
	saveCity(cmd *models.CreateCityCMD) (*models.City, error)
	deleteCity(cityID int64) *models.City
	updateCity(cmd *models.UpdateCityCMD) *models.City
}

type CityStg struct {
	*database.SQLClient
}

func (s *CityStg) saveCity(cmd *models.CreateCityCMD) (*models.City, error) {
	tx, err := s.Begin()

	if err != nil {
		return nil, err
	}

	res, err := tx.Exec(`insert into city (name, latitude, longitude) values (?,?,?)`,
		cmd.Name, cmd.Latitude, cmd.Longitude)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	_ = tx.Commit()

	return &models.City{
		ID:        lastID,
		Name:      cmd.Name,
		Latitude:  cmd.Latitude,
		Longitude: cmd.Longitude,
	}, nil
}

func (s *CityStg) deleteCity(cityID int64) *models.City {
	tx, err := s.Begin()

	if err != nil {
		return nil
	}

	var city models.City
	err = tx.QueryRow(`select name, latitude, longitude from city where id = ?`, cityID).
		Scan(&city.Name, &city.Latitude, &city.Longitude)

	if err != nil {
		_ = tx.Rollback()
		return nil
	}

	_, err = tx.Exec(`delete from city where id = ?`, cityID)

	if err != nil {
		// common error, cannot delete a city with temperatures attached to itself due to constraint limitations.
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()
	return &models.City{
		ID:        cityID,
		Name:      city.Name,
		Latitude:  city.Latitude,
		Longitude: city.Longitude,
	}
}

func (s *CityStg) updateCity(cmd *models.UpdateCityCMD) *models.City {
	tx, err := s.Begin()

	if err != nil {
		return nil
	}

	_, err = tx.Exec(`update city set name = ?, latitude = ?, longitude = ? where id = ?`,
		cmd.Name, cmd.Latitude, cmd.Longitude, cmd.ID)

	if err != nil {
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &models.City{
		ID:        cmd.ID,
		Name:      cmd.Name,
		Latitude:  cmd.Latitude,
		Longitude: cmd.Longitude,
	}
}
