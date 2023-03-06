package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/golang-base/pkg/services"
)

// StatusHandler TODO
// @Summary Shows the status of the application.
// @Description TODO
// @Accept json
// @Produce json
// @Router /status [get]
func StatusHandler(c *gin.Context) {
	status, err := services.GetStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	c.JSON(http.StatusOK, status)
}