package middleware

import (
	"github.com/gin-gonic/gin"
	"jgin/api/model"
)

type GinHandlerDecorator func(gin.HandlerFunc) gin.HandlerFunc

func Decorator(h gin.HandlerFunc, decors ...GinHandlerDecorator) gin.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}

func GetUser(c *gin.Context) (user *model.User) {
	if u, exists := c.Get("CurrentUser"); exists {
		user, _ = u.(*model.User)
	}
	return
}
