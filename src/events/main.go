package main

import (
	"database/sql"
	"fmt"
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

	dbURL, err := get_DB_creds()
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

  events_handler := handlers.NewEventHandler(dbhandler.InitDBHandler(db))

	router := gin.Default()

	router.GET("/manage/health", events_handler.CheckHealth)

	router.GET("/api/v1/events", events_handler.GetAllEventsHandler)
	router.POST("/api/v1/events", events_handler.CreateNewEventHandler)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	
}

func get_DB_creds() (string, error){
	var host, user, dbname, password string
	var port int
 	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)	
	return dbURL, nil
}
