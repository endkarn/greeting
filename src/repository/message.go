package repository

import (
	"greeting/model"

	"github.com/jmoiron/sqlx"
)

type MessageDB interface {
	GetMessage() (model.Greeting, error)
}

type PostgresDB struct {
	DB *sqlx.DB
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
