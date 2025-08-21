package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// setupRouter configures the routes for our application.
// We create it as a separate function so we can also use it in our tests.

func setupRouter () *gin.Engine {
	// gin.Default() creates a new Gin router with some default middleware (for logging, etc.)
	router := gin.Default()


	// This defines a GET endpoint. When a user sends a GET request to "/ping",
	// the function that follows will be executed.
	router.GET("/ping", func(c *gin.Context) {
		// c.JSON() sends a JSON response,
		//http.StatusOK is the standart HTTP code for "OK" (200).
		//gin.h is a shortcut for map[string]interface{}, a generic way to create a JSON object.
		c.JSON(http.StatusOK, gin.H{"message": "pong"})

	})
	return router
}

func main() {
	// set up the router using our function
	router := setupRouter()

	//router.Run() starts the web server on the default port (8080).
	router.Run()
}
