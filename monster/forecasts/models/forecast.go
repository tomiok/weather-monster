package models

type ForecastResponse struct {
	CityID int64   `json:"city_id"`
	Min    float32 `json:"min"`
	Max    float32 `json:"max"`
	Sample int     `json:"sample"`
}
