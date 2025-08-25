package handlers

import (
	"math/rand"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExperimentHandler struct {
	DB *gorm.DB
}


// GetExperiments handles GET /api/experiments
func (h *ExperimentHandler) GetExperiments(c *gin.Context) {
	var experiments []Experiment
	h.DB.Preload("Variations").Find(&experiments)
	c.JSON(http.StatusOK, experiments)
}


func (h *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var input struct {
		Name	string	`json:"name" binding:"required"`
		Variations	[]string `json:"variations" binding:"required,min=2"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experiment := Experiment{Name: input.Name}
	if result := h.DB.Create(&experiment); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create experiment"})
		return
	}

	for _, varName := range input.Variations {
		variation := Variation{Name: varName, ExperimentID: experiment.ID}
		h.DB.Create(&variation)
	}

	var createdExperiment Experiment
	h.DB.Preload("Variations").First(&createdExperiment, experiment.ID)
	c.JSON(http.StatusCreated, createdExperiment)
}

// AssignVariation handles POST /api/experiments/:id/assign
func (h *ExperimentHandler) AssignVariation(c *gin.Context) {
	var experiment Experiment
	if err := h.DB.Preload("Variations").First(&experiment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experiment not found"})
		return
	}

	if len(experiment.Variations) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Experiment has no variations"})
		return
	}

	chosenIndex := rand.Intn(len(experiment.Variations))
	chosenVariation := experiment.Variations[chosenIndex]

	chosenVariation.Participants++
	h.DB.Save((&chosenVariation))

	c.JSON(http.StatusOK, gin.H{"variationName" : chosenVariation.Name, "variationId" : chosenVariation.ID})
}