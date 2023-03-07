package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"

	"github.com/ariel17/be-challenge-arios/pkg/services"
)

func TestStatusHandler(t *testing.T) {
	testCases := []struct {
		name       string
		status     services.Status
		statusCode int
	}{
		{"ok", services.Status{OK: true}, http.StatusOK},
		{"failed", services.Status{OK: false, Detail: "meh"}, http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockStatusService := services.MockStatusService{}
			mockStatusService.On("GetStatus").Return(tc.status)
			statusService = &mockStatusService

			r := gin.Default()
			r.GET(statusPath, StatusHandler)

			req, _ := http.NewRequest(http.MethodGet, statusPath, nil)
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)
			assert.Equal(t, tc.statusCode, rr.Code)
		})
	}
}