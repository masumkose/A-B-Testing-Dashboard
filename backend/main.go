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

	// Group all our API endpoints under /api for better organization
	api := router.Group("/api")
	{
		// Get /api/experiments - List all experiments
		api.GET("/experiments", func(c *gin.Context) {
			var experiments []Experiment
			// Preload("Variations") tells GORM to also fetch the associated variations for each experiment.
			DB.Preload("Variations").Find(&experiments)
			c.JSON(http.StatusOK, experiments)
		})

		api.POST("/experiments", func(c *gin.Context) {
			var input struct {
				Name		string		`json:"name" binding:"required"`
				Variations	[]string	`json:"variations" binding:"required,min=2"`
			}
	
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
				return
			}
	
			experiment := Experiment{Name: input.Name}
			DB.Create(&experiment)
	
	
			// create the associated variations
	
			for _, varName:= range input.Variations {
				variation := Variation{Name: varName, ExperimentID: experiment.ID}
				DB.Create(&variation)
			}
	
			// Fetch the full experiment with variations to treturn it in the response
			var createdExperiment Experiment
	
			DB.Preload("Variations").First(&createdExperiment, experiment.ID)
	
			c.JSON(http.StatusCreated, createdExperiment)
	
		})
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

