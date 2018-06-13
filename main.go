package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/golang-spa/controller"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()

	// files
	router.LoadHTMLFiles("index.html")
	router.Static("/assets", "./assets")

	// routing
	router.GET("/", controller.HomeIndex)
	todo := router.Group("/todo")
	{
		todo.GET("/", controller.TodoIndex)
		todo.POST("/", controller.TodoCreate)
		id := todo.Group("/:id")
		{
			id.GET("/", controller.TodoShow)
			id.PUT("/", controller.TodoUpdate)
			id.DELETE("/", controller.TodoDelete)
		}
	}

	router.Run()
}
