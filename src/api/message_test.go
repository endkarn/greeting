package api_test

import (
	. "greeting/api"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_GreetingHandler_Should_Be_SaWadDeeChaoLok(t *testing.T) {
	expected := `{"message":"สวัสดี ชาวโลก"}`
	request := httptest.NewRequest("GET", "/say", nil)
	writer := httptest.NewRecorder()

	greetingHandler := Greeting{
		DB: &sqlx.DB{},
	}

	testRoute := gin.Default()
	testRoute.GET("/say", greetingHandler.GreetingHandler)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actual))
}
