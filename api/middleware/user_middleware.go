package middleware

import (
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/lib/e"
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
			c.Set("CurrentUid", uid)
			return
		} else {
			common.SetError(c, e.JWT_PARSE_ERROE, nil)
			c.Abort()
			return
		}
	}
}
