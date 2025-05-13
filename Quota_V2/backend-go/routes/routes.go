package routes

import (
	"Quota_V2/backend-go/controllers"
	"Quota_V2/backend-go/utils"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura le rotte dell'applicazione
func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// Rotte protette
	protected := router.Group("/")
	protected.Use(utils.MiddlewareJWT())

	protected.GET("/profile", controllers.GetProfile)
	protected.POST("/logout", controllers.Logout)
}
