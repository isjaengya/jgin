package route

import (
	"github.com/gin-gonic/gin"
	"jgin/api/handler"
	"jgin/api/middleware"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	v1R := r.Group("/v1")

	v1R.Use(middleware.RequestUrlMiddleware())
	v1R.POST("/ping", handler.TestPost)
	v1R.POST("/family", handler.FamilyPost)

	v1R.POST("/user/login", handler.UserLogin)
	//v1R.GET("/user", handler.UserInfo)
	v1R.GET("/user", middleware.Decorator(handler.UserInfo, middleware.VerifyUid))
	//v1R.GET("/inner/user", handler.UserInfo)

	v1R.GET("/user/jwt", handler.CheckUserJwt)

	return r
}
