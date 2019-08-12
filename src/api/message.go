package api

import (
	"greeting/model"
	"greeting/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GreetingHandler(context *gin.Context) {
	greeting, err := repository.GetMessage()
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	message := model.Greeting{
		Message: greeting.Message,
	}
	context.JSON(http.StatusOK, &message)
}
