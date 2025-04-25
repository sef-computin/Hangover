package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
