package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/gateway/service"
)

type ServicesStruct struct {
	EventServiceAddress, 
	UserServiceAddress,
	EnrollmentServiceAddress string

	LogServiceAddress string
}

type GatewayService struct {
	Config ServicesStruct
	// authTokens map[string]time.Time
}

func NewGatewayService(config *ServicesStruct) *GatewayService {
	return &GatewayService{Config: *config}
}

func (gs *GatewayService) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (gs *GatewayService) GetEvents(c *gin.Context){
	_ = c.Request.URL.Query()

	events, err := service.GetEvents(gs.Config.EventServiceAddress)
	if err != nil {
		log.Printf("failed to get response from events service: %v\n", err)
		c.IndentedJSON(http.StatusNoContent, events)
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

func (gs *GatewayService) GetIndexInfo(c *gin.Context) {

}
