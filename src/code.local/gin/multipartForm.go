package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form/autofill.do", func(c *gin.Context) {
		name := c.PostForm("username")
		pwd := c.DefaultPostForm("password", "root")
		c.JSON(http.StatusOK, gin.H{
			"status":   "200",
			"name":     name,
			"password": pwd,
		})
	})
	r.Run(":8082")
}
