package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	cities "github.com/tomiok/weather-monster/monster/cities/web"
	forecasts "github.com/tomiok/weather-monster/monster/forecasts/web"
	temps "github.com/tomiok/weather-monster/monster/temperatures/web"
	whs "github.com/tomiok/weather-monster/monster/webhook/web"
)

// set up all the routes and middleware for the service
func routes(
	city *cities.CityHandler,
	temp *temps.TemperatureHandler,
	webhook *whs.WebhookHandler,
	forecast *forecasts.ForecastHandler,
) *chi.Mux {
	mux := chi.NewMux()

	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	mux.Route("/cities", func(r chi.Router) {
		r.Post("/", city.CreateCityHandler)
		r.Delete("/{cityID:[0-9]+}", city.DeleteCityHandler)
		r.Patch("/{cityID:[0-9]+}", city.UpdateCityHandler)
	})

	mux.Route("/temperatures", func(r chi.Router) {
		r.Post("/", temp.RegisterTemperatureHandler)
	})

	mux.Route("/forecasts", func(r chi.Router) {
		r.Get("/{cityID:[0-9]+}", forecast.GetForecastByCityHandler)
	})

	mux.Route("/webhooks", func(r chi.Router) {
		r.Post("/", webhook.CreateWebhookHandler)
		r.Delete("/{webhookID:[0-9]+}", webhook.DeleteWebhookHandler)
	})

	return mux
}
