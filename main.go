package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	executionTimeHandler "github.com/wahyuhjr-restapi-kpi/handlers/executionTimeHandler"
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
	// Initialize app with dependencies
	app := executionTimeHandler.NewApp(models.DB)

	router := gin.Default()
	router.GET("/ping", CheckDatabase)
	router.GET("/executiontime", app.GetExecutionTime)
	router.GET("/executiontime/:id", app.GetExecutionTimeByID)
	router.POST("/executiontime/create", app.CreateExecutionTime)
	router.DELETE("/executiontime/delete/:id", app.DeleteExecutionTime)
	router.PUT("/executiontime/update/:id", app.UpdateExecutionTime)

	router.Run(":8000")
}
