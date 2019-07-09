package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/lib/e"
	"jgin/api/middleware"
	"jgin/api/schema"
	"jgin/api/util"
)

func UserLogin(c *gin.Context) {
	v := schema.UserLoginSchema{}
	if err := v.Bind(c); err != nil {
		common.SetError(c, e.SHOULD_ERROR, err)
		return
	}

	user, b := schema.VerifyUserLogin(v)
	if b != true {
		common.SetError(c, e.PASSWORD_OR_USERNAME_ERROR, nil)
		return
	}
	jwtLast10 := middleware.GinGetJwt(c, user.Uid)

	go user.SetUserJwtLast10(jwtLast10)
	// user --> json 更新redis user cache
	go user.UpdateRedisCache()
	common.SetOK(c, user)
	return
}

func UserInfo(c *gin.Context) {
	uid := middleware.GetUid(c)
	fmt.Println("当前用户uid为 --> ", uid)
	v := schema.UserQuerySchema{}
	if err := v.Bind(c); err != nil {
		common.SetError(c, e.SHOULD_ERROR, err)
		return
	}
	m := make(map[string]interface{})
	for _, uid := range v.Uids{
		user, err := schema.GetCacheInfoToUser(uid)
		if err == nil {
			m[uid] = user
		}
	}
	common.SetOK(c, m)
	return
}

func CheckUserJwt(c *gin.Context) {
	/*这里不采用装饰器的方式来验证jwt, 验证jwt的接口调用是最多的, 尽量减少里面的逻辑*/
	strict := c.Query("mode") // 是不是严格模式的验证  --> 严格模式get一次redis，否则只验证有没有这个用户
	jwt := c.GetHeader("Authorization")
	uid, ok := util.ParseTokenUid(jwt)
	if ok {
		if strict == "strict" {
			s := util.GetUserJwtLast10(uid)
			if jwt[len(jwt)-10:] != s{
				common.SetError(c, e.JWT_INVALID, nil)
				return
			}
		}
		m := map[string]interface{} {"uid": uid}
		common.SetOK(c, m)
		return
	} else {
		common.SetError(c, e.JWT_PARSE_ERROE, nil)
		return
	}
}