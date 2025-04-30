package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/events/dbhandler"
	"github.com/sef-comp/Hangover/events/models"
)

type EventHandler struct {
	DBHandler dbhandler.EventDB
}

func NewEventHandler(db dbhandler.EventDB) *EventHandler{
  return &EventHandler{DBHandler: db}
}


func (h *EventHandler) GetAllEventsWithParams(c *gin.Context){
	params := make(map[string]any)
	
  for key, values := range c.Request.URL.Query(){
		for _, val := range values{
	 		params[key] = val
		}
  }

	events, err := h.DBHandler.GetEventsWithParams(params)
	if err != nil{
		log.Printf("failed to get events: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, events)
}

func (h *EventHandler) GetAllEventsHandler(c *gin.Context) {
	events, err := h.DBHandler.GetAllEvents()
	if err != nil{
		log.Printf("failed to get events: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, events)
} 

func (h *EventHandler) CreateNewEventHandler(c *gin.Context) {
	var event models.Event

	err := json.NewDecoder(c.Request.Body).Decode(&event)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, nil)
		return
	}

	if err := h.DBHandler.CreateEvent(&event); err != nil {
		log.Printf("Failed to create event: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *EventHandler) DeleteEventHandler(c *gin.Context){
	event_id := c.Param("event_id")
	err := h.DBHandler.DeleteEventByID(event_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (h *EventHandler) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}


