package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"github.com/sef-comp/Hangover/events/dbhandler/mock_dbhandler"
	"github.com/sef-comp/Hangover/events/models"
	"go.uber.org/mock/gomock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllEvents(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockEventDB, record []*models.Event)
	tests := []struct {
		name                 string
		output               []*models.Event
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehaviour: func(r *mock_dbhandler.MockEventDB, output []*models.Event) {
				r.EXPECT().GetAllEvents().Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var events []*models.Event
			dbhandler := mock_dbhandler.NewMockEventDB(c)
			test.mockBehaviour(dbhandler, events)

			handler := EventHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/events", handler.GetAllEventsHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/events", nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestCreateNewEvent(t *testing.T){
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockEventDB, event *bytes.Buffer)
	tests := []struct {
		name                 string
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehaviour: func(r *mock_dbhandler.MockEventDB, event *bytes.Buffer) {
				var event_obj *models.Event = new(models.Event)
			  err := json.Unmarshal(event.Bytes(), event_obj)	
				if err != nil{
				 t.Error(err)
				}
				r.EXPECT().CreateEvent(event_obj)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			valid_event_json, err := json.Marshal(getValidEvent())
			if err != nil{
				t.Error("error creating json body: ", err)
			}
			
			buf := bytes.NewBuffer(valid_event_json)
			

			dbhandler := mock_dbhandler.NewMockEventDB(c)
			test.mockBehaviour(dbhandler, buf)

			handler := EventHandler{DBHandler: dbhandler}

			r := gin.New()
			r.POST("/events", handler.CreateNewEventHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/events", buf)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}

}


func getValidEvent() (ret models.Event){
	ret.EventID, _ = uuid.NewRandom()
	ret.EventName = "Valid Event #1"
	ret.City = "Blaga"
	ret.Description = "Valid Event Data for testing"
	
	ret.StartDt = time.Now().Add(time.Hour*24)
	ret.FinishDt = time.Now().Add(time.Hour*48) 
	ret.CreatedBy, _ = uuid.NewRandom()
	ret.Geolat = -122.4194
	ret.Geolng = 37.7749

	return 
}
