// client package is used to load the client side of the application.
package client

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/config"
)

func Start(r *gin.Engine) {
	// Load the configuration
	Conf := config.LoadConfig()
	// Load the WebPath
	webPath := Conf.WebPath.Dist

	// Load Web
	r.StaticFile("/", webPath+"/index.html")
	r.StaticFile("/favicon.ico", webPath+"/favicon.ico")
	r.StaticFile("/vangoghmuseum-s0164V1962-800.jpg", webPath+"/vangoghmuseum-s0164V1962-800.jpg")
	r.Static("/assets", webPath+"/assets")
}
