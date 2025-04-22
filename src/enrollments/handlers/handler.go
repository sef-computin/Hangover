package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sef-comp/Hangover/enrollments/dbhandler"
)


type EnrollmentHandler struct {
	DBHandler dbhandler.EnrollmentDB
}

func NewEnrollmentHandler(db dbhandler.EnrollmentDB) *EnrollmentHandler{
  return &EnrollmentHandler{DBHandler: db}
}


func (h *EnrollmentHandler) GetAllEnrollments(c *gin.Context){
	events, err := h.DBHandler.GetAllEnrollments()
	if err != nil{
		log.Printf("failed to get events: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, events)
}

func (h *EnrollmentHandler) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}


