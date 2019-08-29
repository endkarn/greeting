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

func (postgres PostgresDB) InsertMessage(newGreeting model.NewGreeting) (model.Greeting, error) {
	greeting := model.Greeting{
		ID:      1,
		Message: "Hello World",
	}
	const query = `INSERT INTO greeting (id, "message") VALUES ($1, $2)`
	tx := postgres.DB.MustBegin()
	tx.MustExec(query, greeting.ID, greeting.Message)
	if err := tx.Commit(); err != nil {
		return model.Greeting{}, err
	}
	return greeting, nil
}

func (postgres PostgresDB) GetMessage() (model.Greeting, error) {
	var greeting model.Greeting
	const query = `SELECT id,message FROM greeting WHERE id = 1`
	err := postgres.DB.Get(&greeting, query)
	if err != nil {
		return model.Greeting{}, err
	}
	return greeting, nil
}
