package models

type Temperature struct {
	ID        int64 `json:"id"`
	CityID    int64 `json:"city_id"`
	Min       int   `json:"min"`
	Max       int   `json:"max"`
	Timestamp int64 `json:"timestamp"`
}

type RegisterTemperatureCMD struct {
	CityID int64 `json:"city_id"`
	Min    int   `json:"min"`
	Max    int   `json:"max"`
}
