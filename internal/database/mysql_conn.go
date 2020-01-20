package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tomiok/weather-monster/internal/logs"
	"go.uber.org/zap"
	"log"
)

type SQLClient struct {
	*sql.DB
}

func NewSQLClient(source string) *SQLClient {

	connection, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err.Error())
	}
	stats := connection.Stats()
	logs.Log().Info("stats: ",
		zap.Int("idle", stats.Idle),
		zap.Int("inUse", stats.InUse),
		zap.Int("openConnections", stats.OpenConnections),
		zap.Int("maxOpenConnections", stats.MaxOpenConnections))
	logs.Log().Info("Connection created")
	return &SQLClient{connection}
}
