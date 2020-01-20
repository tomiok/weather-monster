package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/monster/webhook/models"
)

type WebhookGateway interface {
	RegisterWebhook(cmd *models.CreateWebhookCMD) (*models.Webhook, error)
	DeleteWebhook(id int64) *models.Webhook
}

type WebhookGtw struct {
	WebhookStorage
}

// RegisterWebhook saves a webhook in the DB
func (g *WebhookGtw) RegisterWebhook(cmd *models.CreateWebhookCMD) (*models.Webhook, error) {
	return g.createWebhook(cmd)
}

// DeleteWebhook deletes a webhook from the DB
func (g *WebhookGtw) DeleteWebhook(id int64) *models.Webhook {
	return g.deleteWebhook(id)
}

// NewWebhookGateway easy way to instance a new gateway
func NewWebhookGateway(client *database.SQLClient) WebhookGateway {
	return &WebhookGtw{&WebhookStg{client}}
}
