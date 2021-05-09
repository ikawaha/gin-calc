package main

import (
	"calc/controller"
	"github.com/gin-gonic/gin"
)

import "net/http"

func main() {
	engine:= gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.GET("/add/:x/:y", controller.CalcAdd)
	engine.GET("/div/:x/:y", controller.CalcDiv)
	engine.Run(":8000")
}