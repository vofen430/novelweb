package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	books = map[string]Book{
		"1": {ID: "1", Title: "The Old Man and the Sea", Author: "Ernest Hemingway"},
		"2": {ID: "2", Title: "Moby-Dick", Author: "Herman Melville"},
		"3": {ID: "3", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	}
	users = map[string]User{
		"11111": {Username: "user1", Password: "password1", ID: "11111"},
		"22222": {Username: "user2", Password: "password2", ID: "22222"},
	}
)

type Book struct {
	ID     string `json:"id" form:"ID"`
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

type User struct {
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	ID           string `json:"ID" form:"ID"`
	Avatar       string `json:"avatar" form:"avatar"`
	Nickname     string `json:"nickname" form:"nickname"`
	Introduction string `json:"introduction" form:"introduction"`
	Phone        string `json:"phone" form:"phone"`
	QQ           string `json:"QQ" form:"QQ"`
	Gender       string `json:"gender" form:"gender"`
	Email        string `json:"email" form:"email"`
	Birthday     string `json:"birthday" form:"birthday"`
}

func main() {
	r := gin.Default()

	r.GET("/api/books", handleBooks)
	r.GET("/api/book/:id", handleBook)
	r.POST("/api/register", handleRegister)
	r.GET("/api/user/token", handleLogin)
	r.GET("/api/user/info/:user_id", GetUserInfo)

	fmt.Println("Starting server at port 8000")
	//fmt.Println(users["11111"].ID)
	if err := r.Run(":8000"); err != nil {
		fmt.Println(err)
	}
}
