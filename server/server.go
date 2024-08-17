/*
server package is main package for the server.
*/
package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/pkg/database"

	"github.com/lin-snow/dooo/api"
	// "github.com/lin-snow/dooo/pkg/cors"
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
	// r.Use(cors.SolveCORS())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // 允许的请求源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // 预检请求的缓存时间
	}))

	// API Router
	r = api.Router(r, db)

	// RUN ON PORT
	PORT = "7879"
	r.Run(":" + PORT)
	fmt.Println("Server is running on port: " + PORT)
}
