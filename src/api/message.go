package api

import (
	"greeting/model"
	"greeting/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Greeting struct {
	DB        *sqlx.DB
	MessageDB repository.MessageDB
}

func (api Greeting) GreetingHandler(context *gin.Context) {
	message := model.Greeting{
		Message: "สวัสดีชาวโลก",
	}
	context.JSON(http.StatusOK, &message)
}
