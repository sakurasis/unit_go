package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/user.do", func(c *gin.Context) {
		var form User
		if c.ShouldBind(&form) == nil {
			if form.Name == "root" && form.Password == "root" {
				c.JSON(http.StatusOK, gin.H{"msg": "success"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
			}
		}
	})

	err := r.Run(":8082")
	if err != nil {
		return
	}
}

type User struct {
	Name     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
