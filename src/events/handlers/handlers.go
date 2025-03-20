package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/events/dbhandler"
)

type EventHandler struct {
	DBHandler *dbhandler.EventDB
}

func NewEventHandler(db *dbhandler.EventDB) *EventHandler{
  return &EventHandler{DBHandler: db}
}

func (h *EventHandler) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
