package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hello world"})

	return
}
