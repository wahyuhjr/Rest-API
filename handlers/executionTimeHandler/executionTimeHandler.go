package executionTimeHandler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wahyuhjr-restapi-kpi/db/sqlc"
)

type App struct {
	Queries *sqlc.Queries
}

func NewApp(db *sql.DB) *App {
	return &App{
		Queries: sqlc.New(db),
	}
}

// GetExecutionTime handles the request to get all execution times
func (app *App) GetExecutionTime(c *gin.Context) {
	executionTime, err := app.Queries.GetExecutionTimes(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
			return
		}
		log.Printf("Failed to execute query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"execution_time": executionTime})
}

// GetExecutionTimeByID handles the request to get execution time by ID
func (app *App) GetExecutionTimeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	executionTime, err := app.Queries.GetExecutionTimeByID(context.Background(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
			return
		}
		log.Printf("Failed to execute query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"execution_time": executionTime})
}

// CreateExecutionTime handles the request to create a new execution time
func (app *App) CreateExecutionTime(c *gin.Context) {
	var newExecutionTime sqlc.CreateExecutionTimeParams
	if err := c.ShouldBindJSON(&newExecutionTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	_, err := app.Queries.CreateExecutionTime(context.Background(), newExecutionTime)
	if err != nil {
		log.Printf("Failed to create execution time: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create execution time", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Execution time created successfully", "execution_time": newExecutionTime})
}

// UpdateExecutionTime handles the request to update an execution time by ID
func (app *App) UpdateExecutionTime(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	var updateExecutionTime sqlc.UpdateExecutionTimeParams
	if err := c.ShouldBindJSON(&updateExecutionTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	// Make sure the ID is set in the update params
	updateExecutionTime.ID = int32(id)

	// Perform the update query
	_, err = app.Queries.UpdateExecutionTime(context.Background(), updateExecutionTime)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
			return
		}
		log.Printf("Failed to execute query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Execution time updated successfully", "execution_time": updateExecutionTime})
}

// DeleteExecutionTime handles the request to delete an execution time by ID
func (app *App) DeleteExecutionTime(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	executionTime, err := app.Queries.DeleteExecutionTime(context.Background(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
			return
		}
		log.Printf("Failed to execute query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Execution time deleted successfully", "execution_time": executionTime})
}
