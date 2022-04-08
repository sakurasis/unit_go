package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/jsonp", func(c *gin.Context) {
		m := map[string]interface{}{
			"info": "gulugulu",
		}
		c.JSONP(http.StatusOK, m)
	})
	router.Run(":8081")
}
