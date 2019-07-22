package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/lib/e"
	"jgin/api/model"
	"jgin/api/util"
)

func VerifyUidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")

		uid, ok := util.ParseTokenUid(jwt)
		if ok {
			s := util.GetUserJwtLast10(uid)
			if jwt[len(jwt)-10:] != s {
				common.SetError(c, e.JWT_INVALID, nil)
				c.Abort()
				return
			}
			// 获取用户信息，set一个User类到gin中，在后面结构想要获取当前用户调用：user := middleware.GetUser(c)
			user, err := model.GetCacheInfoToUser(uid)
			if err != nil {
				fmt.Println("middleware get user error: %s", err.Error())
				common.SetError(c, e.MIDDLEWARE_GET_USER_ERROR, err)
				c.Abort()
				return
			}
			c.Set("CurrentUser", user)
			return
		} else {
			common.SetError(c, e.JWT_PARSE_ERROE, nil)
			c.Abort()
			return
		}
	}
}
