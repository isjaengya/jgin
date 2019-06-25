package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tebu_go/api/common"
	"tebu_go/api/lib/e"
	"tebu_go/api/util"
)

type GinHandlerDecorator func(gin.HandlerFunc) gin.HandlerFunc

func Decorator(h gin.HandlerFunc, decors ...GinHandlerDecorator) gin.HandlerFunc {
    for i := range decors {
        d := decors[len(decors)-1-i] // iterate in reverse
        h = d(h)
    }
    return h
}

func VerifyUid(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		jwt := c.GetHeader("Authorization")

		uid, ok := util.ParseTokenUid(jwt)
    	if ok {
    		c.Set("CurrentUid", uid)
			h(c) // 重点！！ 执行初始函数，把这里注释掉就不往下执行了，上面的Decorator貌似没啥作用，fuck
			return
		} else {
			common.SetError(c, e.JWT_PARSE_ERROE, nil)
			return
		}
	}
}

func GetUid(c *gin.Context) (uid string) {
	uid = c.GetString("CurrentUid")
	return
}

func GinGetJwt(c *gin.Context, uid string) (s string) {
	m := map[string]interface{} {"uid": uid}
	jwt := util.CreateJwt(m)
	c.Header("Authorization", jwt)
	fmt.Println("jwt: --> ", jwt)
	return jwt[len(jwt)-10:]
}