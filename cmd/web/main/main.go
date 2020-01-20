package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/tomiok/weather-monster/internal/database"
	"github.com/tomiok/weather-monster/internal/logs"
	cities "github.com/tomiok/weather-monster/monster/cities/web"
	forecasts "github.com/tomiok/weather-monster/monster/forecasts/web"
	temps "github.com/tomiok/weather-monster/monster/temperatures/web"
	whs "github.com/tomiok/weather-monster/monster/webhook/web"
	"os"
)

const (
	migrationsScriptsVersion = 2
	migrationsRootFolder     = "file://migrations"
	mysqlConnStr             = "%s:%s@tcp(%s:3306)/%s?parseTime=true"
)

func main() {
	environment := initEnv()
	logs.InitDefault(environment.env)

	sqlClient := database.NewSQLClient(fmt.Sprintf(mysqlConnStr, environment.dbUser, environment.dbPass, environment.dbUrl,
		environment.dbName))

	doMigrate(sqlClient, environment.dbName)

	mux := routes(
		cities.NewCityHandler(sqlClient),
		temps.NewTemperatureHandler(sqlClient),
		whs.NewWebhookHandler(sqlClient),
		forecasts.NewForecastHandler(sqlClient),
	)

	srv := newServer(environment.port, mux)

	srv.Start()
}

type envSetup struct {
	dbUser string
	dbPass string
	dbUrl  string
	dbName string
	port   string
	env    string
}

func initEnv() *envSetup {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	env := flag.String("E", "local", "the current execution environment")
	flag.Parse()

	if dbPass == "" && dbUser == "" && dbURL == "" && dbName == "" {
		dbUser, dbPass, dbURL, dbName = "root", "root", "localhost", "weather_monster"
	}

	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = "9000"
	}

	return &envSetup{
		dbUser: dbUser,
		dbPass: dbPass,
		dbUrl:  dbURL,
		dbName: dbName,
		port:   serverPort,
		env:    *env,
	}
}

// perform the migration, see migrationsRootFolder constant and migrationsScriptsVersion constant to figure out
// the scripts and which versions are going to be executed
func doMigrate(client *database.SQLClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Sugar().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}
