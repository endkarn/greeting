package api_test

import (
	. "greeting/api"
	"greeting/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_GreetingHandler_Should_Be_SaWadDeeChaoLok(t *testing.T) {
	expected := `{"message":"สวัสดีชาวโลก"}`
	request := httptest.NewRequest("GET", "/say", nil)
	writer := httptest.NewRecorder()

	mockMessageRepository := new(mockGreeting)
	mockMessageRepository.On("GetMessage").Return(model.Greeting{
		ID:      1,
		Message: "สวัสดี ชาวโลก",
	})

	greetingHandler := Greeting{
		DB:        &sqlx.DB{},
		MessageDB: mockMessageRepository,
	}

	testRoute := gin.Default()
	testRoute.GET("/say", greetingHandler.GreetingHandler)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actual))
}
