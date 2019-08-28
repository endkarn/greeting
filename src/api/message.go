package api

import (
	"greeting/model"
	"greeting/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Greeting struct {
	DB        *sqlx.DB
	MessageDB repository.MessageDB
}

func (api Greeting) GreetingHandler(context *gin.Context) {
	greeting, err := api.MessageDB.GetMessage()
	if err != nil {
		log.Println("Handlers GetProductByID error: ", err)
		context.JSON(http.StatusBadRequest, err)
		return
	}
	message := model.Greeting{
		Message: greeting.Message,
	}
	context.JSON(http.StatusOK, &message)
}
