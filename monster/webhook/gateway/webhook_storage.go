package gateway

import (
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	"github.com/tomiok/weather-monster/monster/webhook/models"
)

type WebhookStorage interface {
	createWebhook(cmd *models.CreateWebhookCMD) (*models.Webhook, error)
	deleteWebhook(webhookID int64) *models.Webhook
}

type WebhookStg struct {
	*database.SQLClient
}

func (s *WebhookStg) createWebhook(cmd *models.CreateWebhookCMD) (*models.Webhook, error) {
	tx, err := s.Begin()

	if err != nil {
		return nil, err
	}

	res, err := tx.Exec(`insert into webhook (city_id, callback_url) values (?,?)`,
		cmd.CityID, cmd.CallbackURL)

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

	return &models.Webhook{
		ID:          lastID,
		CityID:      cmd.CityID,
		CallbackURL: cmd.CallbackURL,
	}, nil
}

func (s *WebhookStg) deleteWebhook(webhookID int64) *models.Webhook {
	tx, err := s.Begin()

	if err != nil {
		return nil
	}

	var wh models.Webhook

	err = tx.QueryRow(`select city_id, callback_url from webhook where id = ?`, webhookID).
		Scan(&wh.CityID, &wh.CallbackURL)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}

	_, err = tx.Exec(`delete from webhook where id = ?`, webhookID)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &models.Webhook{
		ID:          webhookID,
		CityID:      wh.CityID,
		CallbackURL: wh.CallbackURL,
	}
}
