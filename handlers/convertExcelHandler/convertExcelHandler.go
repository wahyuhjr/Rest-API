package convertExcelHandler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wahyuhjr-restapi-kpi/db/sqlc"
	"github.com/xuri/excelize/v2"
)

type App struct {
	Queries *sqlc.Queries
}

func NewApp(query *sqlc.Queries) *App {
	return &App{
		Queries: query,
	}
}

// GetExcel handles the request to get all execution times
func (app *App) ConvertExcel(c *gin.Context) {
	executionTimes, err := app.Queries.GetExecutionTimes(context.Background())

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to execute query"})
		return
	}

	//create new excel file
	excelFile := excelize.NewFile()
	sheet := "Sheet1"
	index, _ := excelFile.NewSheet(sheet)

	//set header
	headers := []string{"ID", "Parameter", "Test", "Value", "Deviation"}
	if err := excelFile.SetSheetRow(sheet, "A1", &headers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to set header"})
		return
	}

	//set data
	for i, executionTime := range executionTimes {
		row := []interface{}{
			executionTime.ID,
			executionTime.Parameter.String,
			executionTime.Test.String,
			executionTime.Value.Float64,
			executionTime.Deviation.Float64,
		}
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		if err := excelFile.SetSheetRow(sheet, cell, &row); err != nil {
			log.Printf("Failed to set data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to set data"})
			return
		}

		// set active sheet
		excelFile.SetActiveSheet(index)

		//save file
		savePath := "./data.xlsx"
		if err := excelFile.SaveAs(savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save Excel file"})
			return
		}

		//set headers for excel download
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=data.xlsx")

		//write file to response
		if err := excelFile.Write(c.Writer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to write file"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Excel file created"})
	}
}
