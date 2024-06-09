package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Success": true,
		})
	})

	router.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Success": true,
		})
	})

	router.GET("/show", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Success": true,
		})
	})

	router.PUT("/update", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Success": true,
		})
	})

	router.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Success": true,
		})
	})

	router.Run(":8000")

}
