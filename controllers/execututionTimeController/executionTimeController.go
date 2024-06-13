package executionTimeController

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wahyuhjr-restapi-kpi/db/sqlc"
	"github.com/wahyuhjr-restapi-kpi/models"
)

func GetExecutionTime(c *gin.Context) {
	queries := sqlc.New(models.DB)
	executionTime, err := queries.GetExecutionTimes(context.Background())

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
		}
		log.Fatal("Failed to execute query: ", err)
	}

	c.JSON(http.StatusOK, gin.H{"execution_time": executionTime})
}
