package route

import (
	"github.com/gin-gonic/gin"
	"jgin/api/handler"
	"jgin/api/middleware"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RequestUrlMiddleware())
	// 不需要登录就可以访问的接口
	commonR := r.Group("")
	commonR.POST("/ping", handler.TestPost)
	commonR.POST("/user/login", handler.UserLogin)

	// 需要登录才可以访问的接口
	v1R := r.Group("/v1")
	v1R.Use(middleware.VerifyUidMiddleware())
	v1R.GET("/user/logout", handler.UserLogout)
	v1R.GET("/user", handler.UserInfo)
	v1R.GET("/user/jwt", handler.CheckUserJwt)
	v1R.GET("/user/login/days", handler.GetUserLoginDays)
	v1R.GET("/inner/user", handler.UserInfo)
	//v1R.GET("/user", handler.UserInfo)

	return r
}
