package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	{
		for _, d := range users {
			if d.ID == user.ID && d.Password == user.Password {
				fmt.Println("数据吻合")
				c.JSON(http.StatusOK, gin.H{"message": "数据吻合"})
			}
		}
	}
}
