// +build integration

package repository_test

import (
	"greeting/model"
	. "greeting/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProductDB(t *testing.T) {
	config := Config{
		Host:       "localhost",
		Port:       5432,
		User:       "greeting",
		Password:   "greeting",
		DBName:     "greeting",
		DisableTLS: true,
	}

	db, err := Open(config)
	if err != nil {
		log.Fatal("connect db error: ", err)
	}
	defer db.Close()

	schema := `
	CREATE TABLE greeting (
		id   int,
		message TEXT,
	
		PRIMARY KEY (id)
	);
	`
	resual, err := NewTable(db, schema)
	if err != nil {
		t.Log("ceate tabel error: ", err)

	}
	t.Log(resual)
	postgresDB := PostgresDB{
		DB: db,
	}

	t.Run("Get Message By ID", func(t *testing.T) {
		expected := model.Greeting{
			ID:      1,
			Message: "Hello World",
		}

		actual, err := postgresDB.GetMessage()
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})
	t.Log(resual)
}
