package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
  _ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/events/dbhandler"
	"github.com/sef-comp/Hangover/events/handlers"
)



func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8060"
	}

	dbURL, err := getDatabaseCreds()
	if err != nil{
		panic(err)
	}
	
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to DB")

  events_handler := handlers.NewEventHandler(dbhandler.InitDBHandler(db))

	router := gin.Default()

	router.GET("/manage/health", events_handler.CheckHealth)

	router.GET("/api/v1/events", events_handler.GetAllEventsWithParams)
	router.POST("/api/v1/events", events_handler.CreateNewEventHandler)
	router.DELETE("/api/v1/events/:event_id", events_handler.DeleteEventHandler)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	
}

