package main


import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/masumkose/A-B-Testing-Dashboard/backend/handlers"
)


// setupRouter configures the routes for our application.
// We create it as a separate function so we can also use it in our tests.

func setupRouter () *gin.Engine {
	// gin.Default() creates a new Gin router with some default middleware (for logging, etc.)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	// Create instances of our handlers, passing the database connnection.
	expHandler := &handlers.ExperimentHandler{DB: DB}
	varHandler := &handlers.VariationHandler{DB: DB}
	
	// Group all our API endpoints under /api for better organization
	api := router.Group("/api")
	{
		// Experiment route are now cleanerm just pointing to the handler methods.
		api.GET("/experiments", expHandler.GetExperiments)
		api.POST("/experiments", expHandler.CreateExperiment)
		api.GET("/experiments/:id/assign", expHandler.AssignVariation)

		// Variation routes
		api.POST("/variations/:id/convert", varHandler.ConvertVariation)
	}

	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}) 

	return router
}

func main() {
	ConnectDatabase()
	router := setupRouter()
	router.Run()
}

