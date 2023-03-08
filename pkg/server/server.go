package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/ariel17/be-challenge-arios/api"
	"github.com/ariel17/be-challenge-arios/pkg/clients"
	"github.com/ariel17/be-challenge-arios/pkg/configs"
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
	"github.com/ariel17/be-challenge-arios/pkg/services"
)

const (
	statusPath             = "/status"
	importerPath           = "/importer"
	competitionPlayersPath = "/competitions/:code"
	teamPath               = "/teams/:tla"
	teamPersonsPath        = "/teams/:tla/persons"
	// TODO importer ticket status
)

var (
	statusService   services.StatusService
	importerService services.ImporterService
	playersService  services.FootballService
)

// StartServer creates a new instance of HTTP server with indicated handlers
// configured and begins serving content.
func StartServer() {
	repository := repositories.NewMySQLRepository()
	if err := repository.Connect(); err != nil {
		log.Errorf("cannot connect to database: %v", err.Error())
		panic(err)
	}
	defer repository.Close()

	if err := repository.CreateSchema(); err != nil {
		log.Errorf("cannot create schema: %v", err.Error())
		panic(err)
	}

	apiClient := clients.NewFootballAPIClient()
	statusService = services.NewStatusService(repository)
	importerService = services.NewImporterService(apiClient, repository)
	playersService = services.NewFootballService(repository)

	r := gin.Default()
	r.GET(statusPath, StatusHandler)
	r.POST(importerPath, ImporterHandler)
	r.GET(competitionPlayersPath, PlayersByCompetitionCodeHandler)
	r.GET(teamPath, TeamTLAHandler)
	r.GET(teamPersonsPath, PersonsByTeamTLAHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(fmt.Sprintf(":%d", configs.GetPort())); err != nil {
		panic(err)
	}
}