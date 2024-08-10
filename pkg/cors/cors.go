/*
cors package is a middleware that provides Cross-Origin Resource Sharing (CORS) support for Gin framework.
*/
package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SolveCORS(r *gin.Engine) {
	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // allowed source
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))
}
