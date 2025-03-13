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
	servicesConfig := handlers.ServicesStruct{}

	router := gin.Default()
	gs := handlers.NewGatewayService(&servicesConfig)

	router.GET("/manage/health", gs.CheckHealth)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
