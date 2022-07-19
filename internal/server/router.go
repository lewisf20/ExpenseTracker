package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creates default gin router
func setRouter() *gin.Engine {
	router := gin.Default()

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// Create API route group
	setupApiEndpoints(router)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func setupApiEndpoints(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/register", register)
	api.POST("/signin", signIn)
}
