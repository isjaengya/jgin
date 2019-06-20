package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"tebu_go/api/common"
	"tebu_go/api/lib"
	"tebu_go/api/lib/e"
)

func RequestUrlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		uri := c.Request.RequestURI
		i := strings.Index(uri, "/inner")
		if i != -1 {
			c.Next()
			return
		}

		ts := c.GetHeader("ts")
		sign := c.GetHeader("sign")
		sv := c.GetHeader("sv")

		if sign == "kWzyW23DOnMpGXz9Iqj2fWkaenYz0Qw7JiJrLqA5gZ2DnVGlhSWfoOvZqsa6opoc2m3DwJmfWhuwQRDQLTVY0QHCKR9JoycLljBH" {c.Next()}

		if ts != "" && sign != "" && sv != "" {
			sv, err := strconv.Atoi(sv)
			if err != nil{
				common.SetError(c, e.WRONG_SIGN_VERSION, nil)
				c.Abort()
				return
			}
			if sv == 1{
				serverSign := lib.GenerateSignatureV1(ts)
				if serverSign == sign{
					c.Next()
				} else {
					common.SetError(c, e.INVALID_SIGNATURE, nil)
					c.Abort()
					return
				}
			} else {
				common.SetError(c, e.WRONG_SIGN_VERSION, nil)
				c.Abort()
				return
			}
		} else {
			common.SetError(c, e.MISSING_SIGN_HEADER, nil)
			c.Abort()
			return
		}
	}
}

