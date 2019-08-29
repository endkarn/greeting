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

	resual, err := DropTable(db, "greeting")
	if err != nil {
		t.Log("drop tabel error: ", err)

	}
	t.Log(resual)
}
