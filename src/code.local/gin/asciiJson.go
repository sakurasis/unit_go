package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/json", func(context *gin.Context) {
		data := map[string]interface{}{
			"languages": "GOLANG",
			"feature":   "<B>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})
	err := engine.Run(":8081")
	if err != nil {
		return
	}
}
