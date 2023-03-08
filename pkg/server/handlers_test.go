package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestImporterHandler(t *testing.T) {
	testCases := []struct {
		name              string
		body              string
		importerIsCalled  bool
		importerIsSuccess bool
		status            int
	}{
		{"ok with goroutine ok", `{"code": "abc"}`, true, true, http.StatusCreated},
		{"invalid body", `{"code`, false, false, http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.importerIsCalled {
				mockImporterService := services.MockImporterService{}
				on := mockImporterService.On("ImportDataByCompetitionCode", "abc")
				if tc.importerIsSuccess {
					on.Return(nil)
				} else {
					on.Return(errors.New("mocked error"))
				}
				importerService = &mockImporterService
			}

			r := gin.Default()
			r.POST(importerPath, ImporterHandler)

			req, _ := http.NewRequest(http.MethodPost, importerPath, strings.NewReader(tc.body))
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)
			assert.Equal(t, tc.status, rr.Code)
		})
	}
}