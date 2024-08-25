/*
server package is main package for the server.
*/
package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/api"
	"github.com/lin-snow/dooo/client"
	"github.com/lin-snow/dooo/config"
	"github.com/lin-snow/dooo/pkg/database"
	"github.com/lin-snow/dooo/pkg/solvecors"
)

// PORT is the port number for the server
var PORT string

func Start() {
	// The server will start here

	// Connect to the Database && Migrate the Database
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	database.MigrateDB(db)

	// Default gin router
	r := gin.Default()

	// CORS
	r.Use(solvecors.SolveCORS())

	// API Router
	r = api.Router(r, db)

	// Client
	client.Start(r)

	// RUN ON PORT
	Conf := config.LoadConfig()
	PORT = Conf.Server.Listen
	r.Run(":" + PORT)
	fmt.Println("Server is running on port: " + PORT)
}
