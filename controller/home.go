package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeIndex func
func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
