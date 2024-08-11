/*
server package is main package for the server.
*/
package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/pkg/database"

	"github.com/lin-snow/dooo/api"
	"github.com/lin-snow/dooo/pkg/cors"
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

	// Default
	r := gin.Default()

	// CORS
	r.Use(cors.SolveCORS())

	// API Router
	r = api.Router(r, db)

	// RUN ON PORT
	PORT = "7879"
	r.Run(":" + PORT)
	fmt.Println("Server is running on port: " + PORT)
}
