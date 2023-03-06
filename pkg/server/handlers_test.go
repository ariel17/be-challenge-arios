package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	r := gin.Default()
	r.GET(statusPath, StatusHandler)

	req, _ := http.NewRequest(http.MethodGet, statusPath, nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}