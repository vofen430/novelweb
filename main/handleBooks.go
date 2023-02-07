package main

import (
	"github.com/gin-gonic/gin"
)

func handleBooks(c *gin.Context) {
	var booksList = []Book{}
	for _, b := range books {
		booksList = append(booksList, b)
	}
	c.Header("Content-Type", "application/json")
	c.JSON(200, booksList)
}
