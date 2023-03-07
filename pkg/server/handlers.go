package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/be-challenge-arios/pkg/services"
)

// StatusHandler TODO
// @Summary Shows the status of the application.
// @Description TODO
// @Accept json
// @Produce json
// @Router /status [get]
func StatusHandler(c *gin.Context) {
	status := services.GetStatus()
	if status.IsError() {
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	c.JSON(http.StatusOK, status)
}