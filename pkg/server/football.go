package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/be-challenge-arios/pkg/models"
)

const (
	codeParamName        = "code"
	teamNameQueryName    = "teamName"
	tlaParamName         = "tla"
	showPlayersQueryName = "showPlayers"
)

type PlayersResult struct {
	Status
	Players []models.Person `json:"players,omitempty"`
}

// PlayersByCompetitionCodeHandler Serves all players participating on a given
// competition.
// @Summary Shows all players from a given competition.
// @Description Given the competition code, if it exists on database, returns all players from all participating teams.
// @Param code path string true "Competition code to filter players."
// @Param teamName query string false "Team name to filter players by"
// @Produce json
// @Success 200 {object} PlayersResult
// @Failure 400 {object} PlayersResult
// @Failure 500 {object} PlayersResult
// @Router /competitions/:code/players [get]
func PlayersByCompetitionCodeHandler(c *gin.Context) {
	code := c.Param(codeParamName)
	teamName, _ := c.GetQuery(teamNameQueryName)
	players, competitionExists, err := playersService.GetPlayersByCompetitionCode(code, teamName)
	if err != nil {
		result := PlayersResult{
			Status: Status{
				Detail: err.Error(),
			},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	if !competitionExists {
		result := PlayersResult{
			Status: Status{
				Detail: fmt.Sprintf("competition %s does not exist on database", code),
			},
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result := PlayersResult{
		Status: Status{
			OK: true,
		},
		Players: players,
	}
	c.JSON(http.StatusOK, result)
}

type TeamResult struct {
	Status
	Team *models.Team `json:"team,omitempty"`
}

// TeamByTLAHandler Retrieves details for indicated team. Can add players/coach
// data if specified.
// @Summary Retrieves indicated team details with players/coach.
// @Description If indicated players/coach also can be resolved if they exist.
// @Param tla path string true "Team TLA value to fetch."
// @Param showPlayers query bool false "Resolve team players/coach if present."
// @Produce json
// @Success 200 {object} TeamResult
// @Failure 404 {object} TeamResult
// @Failure 500 {object} TeamResult
// @Router /teams/:tla [get]
func TeamByTLAHandler(c *gin.Context) {
	tla := c.Param(tlaParamName)
	rawShowPlayers, _ := c.GetQuery(showPlayersQueryName)
	showPlayers, _ := strconv.ParseBool(rawShowPlayers)
	team, err := playersService.GetTeamByTLA(tla, showPlayers)
	if err != nil {
		result := TeamResult{
			Status: Status{
				Detail: err.Error(),
			},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	if team == nil {
		result := TeamResult{
			Status: Status{
				Detail: fmt.Sprintf("team %s does not exist on database", tla),
			},
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	result := TeamResult{
		Status: Status{
			OK: true,
		},
		Team: team,
	}
	c.JSON(http.StatusOK, result)
}

type PersonsResult struct {
	Status
	Persons []models.Person `json:"persons"`
}

// PersonsByTeamTLAHandler Retrieves all persons on a given team.
// @Summary Retrieves all persons on a team (players/coach).
// @Description Retrieves all persons on a team (players/coach).
// @Param tla path string true "Team TLA value to fetch."
// @Produce json
// @Success 200 {object} PersonsResult
// @Failure 500 {object} PersonsResult
// @Router /teams/:tla/persons [get]
func PersonsByTeamTLAHandler(c *gin.Context) {
	tla := c.Param(tlaParamName)
	persons, err := playersService.GetPersonsByTeamTLA(tla)
	if err != nil {
		result := PersonsResult{
			Status: Status{
				Detail: err.Error(),
			},
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result := PersonsResult{
		Status: Status{
			OK: true,
		},
		Persons: persons,
	}
	c.JSON(http.StatusOK, result)
}