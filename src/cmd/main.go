package main

import (
	"flag"
	"fmt"
	"greeting/api"
	"greeting/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var port, dbhost, dbschema, dbusername, dbpassword, disableTLS string
	var dbport int

	flag.StringVar(&port, "port", "3000", "port for open service")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database host name")
	flag.IntVar(&dbport, "dbport", 5432, "database port")
	flag.StringVar(&dbschema, "dbschema", "smalldoc", "database schema name")
	flag.StringVar(&dbusername, "dbusername", "smalldoc", "database user name")
	flag.StringVar(&dbpassword, "dbpassword", "example", "database password")
	flag.StringVar(&disableTLS, "disableTLS", "Y", "database disableTLS[Y/n]")
	flag.Parse()

	var databaseTSL bool
	if disableTLS == "n" {
		databaseTSL = false
	} else {
		databaseTSL = true
	}

	dbConfig := repository.Config{
		User:       dbusername,
		Password:   dbpassword,
		Host:       dbhost,
		Port:       dbport,
		Name:       dbschema,
		DisableTLS: databaseTSL,
	}

	db, err := repository.Open(dbConfig)
	if err != nil {
		log.Fatal("connecting database fail", err)
	}

	engine := gin.Default()
	greetingHandler := api.Greeting{
		DB: db,
	}

	engine.GET("/say/", greetingHandler.GreetingHandler)
	log.Fatal(engine.Run(fmt.Sprintf(":%s", port)))
}
