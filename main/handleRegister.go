package main

import (
	"github.com/gin-gonic/gin"
)

func handleRegister(c *gin.Context) {
	var newUser User
	var found = true
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	{
		for _, d := range users {
			if d.ID == newUser.ID {
				c.AbortWithStatusJSON(400, gin.H{"error": "id重复"})
				found = false
				break
			}
		}
	}
	switch {
	case newUser.ID == "":
		c.AbortWithStatusJSON(400, gin.H{"error": "error"})
		return
	case newUser.Username == "":
		c.AbortWithStatusJSON(400, gin.H{"error": "error"})
		return
	case newUser.Password == "":
		c.AbortWithStatusJSON(400, gin.H{"error": "error"})
		return
	case found == true:
		key := newUser.ID
		users[key] = newUser
		c.JSON(200, gin.H{"message": "User created successfully"})
	default:
		return
	}
}
