package main

import (
	"backend/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// Respond with "Hello, World!"
		c.String(200, "Hello, World 23!\n")
	})

	router.GET("/organization", func(c *gin.Context) {
		services.GetOrganizations()
		c.JSON(http.StatusOK, gin.H{"prop": "value"})
	})

	router.GET("/organization/:id", func(c *gin.Context) {
		// Get the id parameter from the route
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	router.POST("/organization", func(c *gin.Context) {
		services.GetOrganizations()
		c.String(200, "Getting Organizations\n")
	})

	router.Run("localhost:8080")
}
