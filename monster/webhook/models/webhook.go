package models

type Webhook struct {
	ID          int64  `json:"id"`
	CityID      int64  `json:"city_id"`
	CallbackURL string `json:"callback_url"`
}

type CreateWebhookCMD struct {
	CityID      int64  `json:"city_id"`
	CallbackURL string `json:"callback_url"`
}
