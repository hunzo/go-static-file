package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Login",
		})
	})

	r.POST("/authen", func(c *gin.Context) {

		type UserLogin struct {
			Username string
			Password string
		}
		var u UserLogin

		u.Username = c.PostForm("username")
		u.Password = c.PostForm("password")

		fmt.Println(u)

		c.JSON(http.StatusOK, gin.H{
			"status":    "getPostForm",
			"UserLogin": u,
		})
	})

	r.Run()
}
