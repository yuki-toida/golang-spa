package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yuki-toida/golang-spa/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", controller.HomeIndex)
	todo := router.Group("/todo")
	{
		todo.GET("/", controller.TodoIndex)
		todo.POST("/", controller.TodoCreate)
		todo.GET("/:id", controller.TodoShow)
	}

	router.Run()
}
