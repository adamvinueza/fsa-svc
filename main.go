package main

import (
	"github.com/adamvinueza/fsa"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = fsa.DFA{}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
