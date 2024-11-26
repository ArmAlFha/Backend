package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Employee API Method
func main() {
	router := gin.Default()
	router.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Fish",
		})
	})
	//Default API Method
	router.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Tiger",
		})
	})
	router.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "CAT",
		})
	})
	router.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DOG",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
