package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServicesStruct struct {
}

type GatewayService struct {
	Config ServicesStruct
	// authTokens map[string]time.Time
}

func NewGatewayService(config *ServicesStruct) *GatewayService {
	return new(GatewayService)
}

func (gs *GatewayService) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (gs *GatewayService) GetIndexInfo(c *gin.Context) {

}
