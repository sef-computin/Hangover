package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/events/dbhandler"
	"github.com/sef-comp/Hangover/events/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// r := handlers.Router()
  db_configuration := dbhandler.DBConfig{}

  eh := handlers.NewEventHandler(dbhandler.NewDBHandler(db_configuration))

	router := gin.Default()

	router.GET("/manage/health", eh.CheckHealth)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
