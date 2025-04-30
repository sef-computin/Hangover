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
	event_id, err := uuid.Parse(c.Param("event_id"))
	if err != nil{
		c.Status(http.StatusBadRequest)
	}
	user_id, err := uuid.Parse(c.Param("user_id"))
	if err != nil{
		c.Status(http.StatusBadRequest)
	}
	enrollment_id, err := h.DBHandler.Enroll(user_id, event_id)
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Data(http.StatusOK, "text/plain", []byte(enrollment_id.String()))
}

func (h *EnrollmentHandler) CancelEnroll(c *gin.Context){
	enrollment_id, err := uuid.Parse(c.Param("enrollment_id"))
	if err != nil{
		c.Status(http.StatusBadRequest)
	}
	err = h.DBHandler.CancelEnroll(enrollment_id)
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *EnrollmentHandler) CheckHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

