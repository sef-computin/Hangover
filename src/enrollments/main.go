package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/enrollments/dbhandler"
	"github.com/sef-comp/Hangover/enrollments/handlers"
)


func main(){
	
	port := os.Getenv("PORT")

	if port == "" {
		port = "8061"
	}

	db_url, err := getDatabaseCreds()
	if err != nil{
		panic(err)
	}
	
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to DB")

  enrollments_handler := handlers.NewEnrollmentHandler(dbhandler.InitDBHandler(db))

	router := gin.Default()

	router.GET("/manage/health", enrollments_handler.CheckHealth)

	router.GET("/api/v1/enrollments", enrollments_handler.GetAllEnrollments)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	
}
