package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/ariel17/be-challenge-arios/api"
	"github.com/ariel17/be-challenge-arios/pkg/configs"
)

const statusPath = "/status"

// StartServer creates a new instance of HTTP server with indicated handlers
// configured and begins serving content.
func StartServer() {
	r := gin.Default()
	r.GET(statusPath, StatusHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(fmt.Sprintf(":%d", configs.GetPort())); err != nil {
		panic(err)
	}
}