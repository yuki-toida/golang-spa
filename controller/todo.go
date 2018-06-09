package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/golang-spa/model/todo"
)

// TodoIndex func
func TodoIndex(c *gin.Context) {
	c.JSON(http.StatusOK, todo.All())
}

// TodoCreate func
func TodoCreate(c *gin.Context) {
	params := &struct {
		Name string `json:"name"`
	}{}
	c.ShouldBindJSON(&params)
	fmt.Printf("params is %v", params)

	c.JSON(http.StatusOK, todo.Create(params.Name))
}

// TodoShow func
func TodoShow(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, todo.Get(id))
}
