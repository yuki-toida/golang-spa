package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/golang-spa/model/todo"
)

// TodoIndex func
func TodoIndex(c *gin.Context) {
	c.JSON(http.StatusOK, todo.All())
}

// TodoShow func
func TodoShow(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, todo.Get(id))
}

// TodoCreate func
func TodoCreate(c *gin.Context) {
	params := &struct {
		Name string `json:"name"`
	}{}
	c.ShouldBindJSON(&params)
	c.JSON(http.StatusOK, todo.Create(params.Name))
}

// TodoUpdate func
func TodoUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	params := &struct {
		Name string `json:"name"`
	}{}
	c.ShouldBindJSON(&params)
	c.JSON(http.StatusOK, todo.Update(id, params.Name))
}

// TodoDelete func
func TodoDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, todo.Delete(id))
}
