package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *EnrollmentHandler) Enroll(c *gin.Context){
	eventID, err := uuid.Parse(c.Param("eventID"))
	if err != nil{
		c.Status(http.StatusBadRequest)
	}
	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil{
		c.Status(http.StatusBadRequest)
	}
	enrollment_id, err := h.DBHandler.Enroll(userID, eventID)
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Data(http.StatusOK, "text/plain", []byte(enrollment_id.String()))
}

func (h *EnrollmentHandler) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

