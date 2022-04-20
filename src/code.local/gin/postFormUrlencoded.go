package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/info", func(c *gin.Context) {
		no := c.Query("no")
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")
		username := c.PostForm("username")
		fmt.Printf("ID: %s\t page: %s\t limit: %s\t username: %s\t\n", no, page, limit, username)

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "操作成功",
		})
	})
	r.Run(":8082")
}
