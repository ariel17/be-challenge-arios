package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ImporterCommand struct {
	Code string `json:"code"`
}

// ImporterHandler Endpoint handler to enqueue a background process collecting
// data from football-data.org API.
// @Summary Imports football data by competition code.
// @Description Enqueues data scrapping from football-data.org API based on competition code. It is a background process so this endpoint only reflects the state of the petition.
// @Accept json
// @Param message body ImporterCommand true "Competition code to import."
// @Produce json
// @Success 201 {object} Status
// @Failure 400 {object} Status
// @Failure 500 {object} Status
// @Router /importer [post]
func ImporterHandler(c *gin.Context) {
	var body ImporterCommand
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, Status{Detail: err.Error()})
		return
	}

	if body.Code == "" {
		c.JSON(http.StatusBadRequest, Status{Detail: "missing code"})
		return
	}

	go func() {
		err := importerService.ImportDataByCompetitionCode(body.Code)
		if err != nil {
			log.Errorf("failed to complete data import for %s %v", body.Code, err)
			return
		}
		log.Infof("data import completed for %s %v", body.Code, err)
	}()

	c.JSON(http.StatusCreated, Status{OK: true, Detail: "queued"})
}

// TODO handler to check importer ticket status