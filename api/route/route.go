package route

import (
	"github.com/gin-gonic/gin"
	"tebu_go/api/handler"
	"tebu_go/api/middleware"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	v1R := r.Group("/v1")
	v1R.Use(middleware.RequestUrlMiddleware())
	v1R.POST("/ping", middleware.Decorator(handler.TestPost, middleware.VerifyUid))
	v1R.POST("/family", handler.FamilyPost)

	return r
}
