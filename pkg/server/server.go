package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/ariel17/be-challenge-arios/api"
	"github.com/ariel17/be-challenge-arios/pkg/clients"
	"github.com/ariel17/be-challenge-arios/pkg/configs"
	"github.com/ariel17/be-challenge-arios/pkg/repositories"
	"github.com/ariel17/be-challenge-arios/pkg/services"
)

const (
	statusPath = "/status"
	// TODO importer
	// TODO importer ticket status
	// TODO players by league code
	// TODO team by tla
	// TODO players by team tla
)

var (
	statusService   services.StatusService
	importerService services.ImporterService
)

// StartServer creates a new instance of HTTP server with indicated handlers
// configured and begins serving content.
func StartServer() {
	repository := repositories.NewMySQLRepository()
	apiClient := clients.NewFootballAPIClient()
	statusService = services.NewStatusService(repository)
	importerService = services.NewImporterService(apiClient, repository)

	r := gin.Default()
	r.GET(statusPath, StatusHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(fmt.Sprintf(":%d", configs.GetPort())); err != nil {
		panic(err)
	}
}