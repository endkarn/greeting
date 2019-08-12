package repository

import (
	"greeting/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetMessage_Should_Be_ID_1_Message_HelloWorld(t *testing.T) {
	dbConfig := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "greeting",
		Password: "greeting",
		SSLMode:  "disable",
		DBName:   "greeting",
	}
	connectionDB, _ := dbConfig.Open()
	db := PostgresDB{
		DB: connectionDB,
	}
	expected := model.Greeting{
		ID:      1,
		Message: "HelloWorld",
	}

	actual, _ := db.GetMessage()

	assert.Equal(t, expected, actual)

}
