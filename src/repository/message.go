package repository

import (
	"fmt"
	"greeting/model"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "greeting"
	password = "greeting"
	dbname   = "greeting"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	SSLMode  string
	DBName   string
}

type PostgresDB struct {
	DB *sqlx.DB
}

func (cfg Config) Open() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	return sqlx.Connect("postgress", psqlInfo)
}

func (postgress PostgresDB) GetMessage() (model.Greeting, error) {
	var greeting model.Greeting
	const query = `SELECT id,message FROM greeting WHERE id = 1`
	err := postgress.DB.Get(&greeting, query)
	if err != nil {
		return model.Greeting{}, err
	}
	return greeting, nil
}
