package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VariationHandler struct {
	DB *gorm.DB
}

func (h *VariationHandler) ConvertVariation(c *gin.Context) {
	var variation Variation
	if err := h.DB.First(&variation, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Variation not found"})
		return
	}

	variation.Conversions++
	h.DB.Save(&variation)

	c.JSON(http.StatusOK, gin.H{"message": "Conversion recorded"})
}