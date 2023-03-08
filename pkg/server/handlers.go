package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// StatusHandler TODO
// @Summary Shows the status of the application.
// @Description TODO
// @Accept json
// @Produce json
// @Router /status [get]
func StatusHandler(c *gin.Context) {
	status := statusService.GetStatus()
	if status.OK {
		c.JSON(http.StatusOK, status)
		return
	}
	c.JSON(http.StatusInternalServerError, status)
}

type ImporterCommand struct {
	Code string `json:"code"`
}

type ImporterResult struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func ImporterHandler(c *gin.Context) {
	var body ImporterCommand
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, ImporterResult{Status: "error", Detail: err.Error()})
		return
	}

	if body.Code == "" {
		c.JSON(http.StatusBadRequest, ImporterResult{Status: "rejected", Detail: "missing code"})
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

	c.JSON(http.StatusCreated, ImporterResult{Status: "queued"})
}

func PlayersByCompetitionCodeHandler(c *gin.Context) {
	code := c.Param("code")
	teamName, _ := c.GetQuery("teamName")
	players, competitionExists, err := playersService.GetPlayersByCompetitionCode(code, teamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !competitionExists {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("competition %s does not exist on database", code))
		return
	}
	c.JSON(http.StatusOK, players)
}

func TeamTLAHandler(c *gin.Context) {
	tla := c.Param("tla")
	rawShowPlayers, _ := c.GetQuery("showPlayers")
	showPlayers, _ := strconv.ParseBool(rawShowPlayers)
	team, err := playersService.GetTeamByTLA(tla, showPlayers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if team == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("team %s does not exist on database", tla))
		return
	}
	c.JSON(http.StatusOK, team)
}

func PersonsByTeamTLAHandler(c *gin.Context) {
	tla := c.Param("tla")
	persons, err := playersService.GetPersonsByTeamTLA(tla)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, persons)
}