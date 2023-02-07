package main

import "github.com/gin-gonic/gin"

func handleBook(c *gin.Context) {
	bookID := c.Param("id")
	b, ok := books[bookID]
	if !ok {
		c.AbortWithStatus(404)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(200, b)
}
