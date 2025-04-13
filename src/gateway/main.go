package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/gateway/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// r := handlers.Router()
	servicesConfig := handlers.ServicesStruct{
		EventServiceAddress: "http://event-service:8080",
		// EventServiceAddress: "http://localhost:8060",
		EnrollmentServiceAddress: "http://enrollment-service:8080",
		UserServiceAddress: "http://user-service:8080",
		LogServiceAddress: "http://log-service:8070",
	}

	router := gin.Default()
	gs := handlers.NewGatewayService(&servicesConfig)

	router.GET("/manage/health", gs.CheckHealth)
	router.GET("/api/v1/events", gs.GetEvents)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
