package routes

import (
	"Quota_V2/backend-go/controllers"
	"Quota_V2/backend-go/utils"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura le rotte dell'applicazione
func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)

	// Rotte protette
	protected := router.Group("/")
	protected.Use(utils.MiddlewareJWT())

	protected.GET("/profile", func(c *gin.Context) {
		userID := c.GetString("user_id")
		c.JSON(200, gin.H{"message": "Profilo utente", "user_id": userID})
	})
}
