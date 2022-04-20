package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	flag := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if flag != nil {
		return
	}
	engine.LoadHTMLGlob("templates/**/*")
	engine.GET("/info", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sys/left.tmpl", gin.H{
			"title": "HOME PAGE",
		})
	})
	engine.GET("/right", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/right.tmpl", gin.H{
			"title": "list PAGE",
		})
	})
	err := engine.Run(":8081")
	if err != nil {
		return
	}
}
