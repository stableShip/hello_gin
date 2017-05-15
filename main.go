package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Print("test")
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}