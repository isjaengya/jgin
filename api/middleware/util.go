package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jgin/api/model"
	"jgin/api/util"
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

func GinGetJwt(c *gin.Context, uid string) (s string) {
	m := map[string]interface{}{"uid": uid}
	jwt := util.CreateJwt(m)
	c.Header("Authorization", jwt)
	fmt.Println("jwt: --> ", jwt)
	return jwt[len(jwt)-10:]
}
