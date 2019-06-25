package handler

import (
	"fmt"
	"tebu_go/api/common"
	"tebu_go/api/lib/e"
	//"github.com/francoispqt/gojay"
	"github.com/gin-gonic/gin"

	"tebu_go/api/middleware"
	"tebu_go/api/model"
	"tebu_go/api/schema"
)

func UserLogin(c *gin.Context) {
	var userLoginSchema schema.UserLoginSchema
	if err := c.ShouldBind(&userLoginSchema); err != nil {
		fmt.Println(userLoginSchema)
		common.SetError(c, e.PARAM_ERROR, err)
		return
	}
	fmt.Println(userLoginSchema)

	user, b := model.VerifyUserLogin(userLoginSchema)
	if b != true {
		fmt.Println("用户验证失败，重新尝试")
		return
	}
	fmt.Println(b, "bool", user, "user")
	jwtLast10 := middleware.GinGetJwt(c, user.Uid)


	go user.SetUserJwtLast10(jwtLast10)
	// user --> json 更新redis user cache
	go user.UpdateRedisCache()

	common.SetOK(c, user)
	return
}

func UserInfo(c *gin.Context) {
	uid := middleware.GetUid(c)
	fmt.Println("uid", uid)
	user := model.GetRedisInfoToUser(uid)
	fmt.Println(user)
	common.SetOK(c, user)
	return
}