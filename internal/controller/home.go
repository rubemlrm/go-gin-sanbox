package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController ...
type HomeController struct{}

//Get ...
func (ctrl HomeController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}
