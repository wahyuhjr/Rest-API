package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	executionTimeController "github.com/wahyuhjr-restapi-kpi/controllers/execututionTimeController"
	"github.com/wahyuhjr-restapi-kpi/models"
)

func CheckDatabase(c *gin.Context) {
	err := models.DB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database is not connected",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Database is connected",
	})
}

func main() {
	models.ConnectDatabase()

	router := gin.Default()
	router.GET("/ping", CheckDatabase)
	router.GET("/executiontime", executionTimeController.GetExecutionTime)

	router.Run(":8000")
}
