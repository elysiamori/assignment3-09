package requestapi

import (
	"fmt"
	"net/http"

	"github.com/elysiamori/assignment3-09/config"
	"github.com/elysiamori/assignment3-09/models"
	"github.com/gin-gonic/gin"
)

func UpdateWeatherData(c *gin.Context) {
	db := config.DBConn()
	var dataToUpdate models.Weather
	err := c.ShouldBindJSON(&dataToUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error: Binding Data", err.Error())
		return
	}

	errUp := db.Model(&models.Weather{}).Where("id = ?", 1).Updates(&dataToUpdate).Error

	if errUp != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data in database"})
		fmt.Println("Failed updating data in database:", errUp)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been updated"})
}
