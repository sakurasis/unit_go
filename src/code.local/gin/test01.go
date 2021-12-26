package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user.do", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "success",
			"data": "user info",
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
