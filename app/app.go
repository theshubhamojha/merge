package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var configuration Configurations
var db *sqlx.DB
var dependencies *Dependencies

type Configurations struct {
	APP_PORT             int
	APP_HOST             string
	PG_DATABASE_HOST     string
	PG_DATABASE_PORT     string
	PG_DATABASE_USER_ID  string
	PG_DATABASE_PASSWORD string
	PG_DB_NAME           string
	JWT_EXPIRY           int
	JWT_SECRET           string
}

// initaliseConfiguration will initialse the local environment from the file `config.yml`
func initaliseConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("error reading config file, error: %s", err.Error()))
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Sprintf("unable to decode configurations, error: %s", err.Error()))
	}
}

func initaliseDatabase() {
	database, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			configuration.PG_DATABASE_HOST,
			configuration.PG_DATABASE_PORT,
			configuration.PG_DATABASE_USER_ID,
			configuration.PG_DB_NAME,
		))
	if err != nil {
		panic(fmt.Sprintf("failed to initalise postgres database, error: %s", err.Error()))
	}

	fmt.Println("successfully connected to the database")
	db = database
}

func InitaliseApp() {
	initaliseConfiguration()
	initaliseDatabase()
	dependencies = initialiseDependencies()
}

func GetDB() *sqlx.DB {
	return db
}

func GetConfiguration() Configurations {
	return configuration
}

func GetDependencies() *Dependencies {
	return dependencies
}
