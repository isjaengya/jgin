package middleware

import (
	"github.com/gin-gonic/gin"
	"tebu_go/api/common"
	"tebu_go/api/lib"
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

		key := lib.JwtKey
		jwt := c.GetHeader("jwt")
		uid, ok := util.ParseTokenUid(jwt, key)
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
