package handler

import (
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/lib/e"
	"jgin/api/middleware"
	"jgin/api/model"
	"jgin/api/schema"
)

func UserLogin(c *gin.Context) {
	v := schema.UserLoginSchema{}
	if err := v.Bind(c); err != nil {
		common.SetError(c, e.SHOULD_ERROR, err)
		return
	}

	user, b := model.VerifyUserLogin(v)
	if b != true {
		common.SetError(c, e.PASSWORD_OR_USERNAME_ERROR, nil)
		return
	}
	jwtLast10 := middleware.GinGetJwt(c, user.Uid)

	go user.SetUserJwtLast10(jwtLast10)
	//go tasks.AsyncHelloWorld(user.CreateAt)

	common.SetOK(c, user)
	return
}

func UserLogout(c *gin.Context) {
	user := middleware.GetUser(c)
	go user.DeleteUserJwtLast10()
	c.Header("Authorization", "")
	common.SetOK(c, "ok")
	return
}

func UserInfo(c *gin.Context) {
	v := schema.UserQuerySchema{}
	if err := v.Bind(c); err != nil {
		common.SetError(c, e.SHOULD_ERROR, err)
		return
	}
	m := make(map[string]interface{})
	for _, uid := range v.Uids {
		user, err := model.GetCacheInfoToUser(uid)
		if err == nil {
			m[uid] = user
		}
	}
	common.SetOK(c, m)
	return
}

func CheckUserJwt(c *gin.Context) {
	user := middleware.GetUser(c)
	common.SetOK(c, user)
	return
}
