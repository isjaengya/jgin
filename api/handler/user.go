package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
	"tebu_go/api/common"
	"tebu_go/api/lib/e"
	"tebu_go/api/middleware"
	"tebu_go/api/model"
	"tebu_go/api/schema"
	"tebu_go/api/util"
)

func UserLogin(c *gin.Context) {
	var userLoginSchema schema.UserLoginSchema
	if err := c.ShouldBind(&userLoginSchema); err != nil {
		common.SetError(c, e.PARAM_ERROR, err)
		return
	}

	user, b := model.VerifyUserLogin(userLoginSchema)
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
	uidS := c.Query("uid")
	if uidS == "" {
		common.SetError(c, e.DATA_ERROE, nil)
		return
	}
	i := strings.Index(uidS, ",")
	if i == -1 {
		user, err := model.GetCacheInfoToUser(uidS)
		if err != nil{
			common.SetError(c, e.DATA_ERROE, err)
			return
		}
		common.SetOK(c, user)
		return
	} else {
		// 传递多个uid
		m := make(map[string]interface{})
		uidL := util.SplitUid(uidS)
		for _, uid := range uidL{
			user, err := model.GetCacheInfoToUser(uid)
			if err == nil {
				m[uid] = user
			}
		}
		common.SetOK(c, m)
		return
	}
}

func CheckUserJwt(c *gin.Context) {
	/*这里不采用装饰器的方式来验证jwt, 验证jwt的接口调用是最多的, 尽量减少里面的逻辑*/
	strict := c.Query("mode") // 是不是严格模式的验证  --> 严格模式get一次redis，否则只验证有没有这个用户
	jwt := c.GetHeader("Authorization")
	uid, ok := util.ParseTokenUid(jwt)
	if ok {
		if strict == "strict" {
			s := model.GetUserJwtLast10(uid)
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