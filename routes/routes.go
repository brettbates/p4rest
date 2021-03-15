package routes

import (
	"github.com/brettbates/p4rest/handlers"
	"github.com/gin-gonic/gin"
)

// Initialize adds all our routes to the router
func Initialize(router *gin.Engine) {

	// Handle the index route
	router.GET("/fixes/", handlers.Fixes)
}
