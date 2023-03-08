package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	OK     bool   `json:"ok"`
	Detail string `json:"detail"`
}

// StatusHandler Shows application health state.
// @Summary Shows the status of the application.
// @Description Returns a JSON reflecting the application's health.
// @Produce json
// @Success 200 {object} Status
// @Failure 500 {object} Status
// @Router /status [get]
func StatusHandler(c *gin.Context) {
	if err := statusService.GetStatus(); err != nil {
		c.JSON(http.StatusInternalServerError, Status{Detail: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Status{OK: true})
}