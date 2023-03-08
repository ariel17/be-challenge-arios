package main

import (
	"github.com/ariel17/be-challenge-arios/pkg/server"
)

// @title           BE Challenge by Ariel Gerardo Ríos
// @version         0.1
// @description     A challenge that uses football-data.org data on its own models.

// @contact.name   Ariel Gerardo Ríos
// @contact.url    http://ariel17.com.ar
// @contact.email  arielgerardorios@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	server.StartServer()
}