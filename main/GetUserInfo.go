package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	userID := c.Param("user_id")
	user, ok := users[userID]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
