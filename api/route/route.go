package route

import (
	"github.com/gin-gonic/gin"
	"tebu_go/api/handler"
	"tebu_go/api/middleware"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	v1_r := r.Group("/v1")
	v1_r.Use(middleware.RequestUrlMiddleware())
	v1_r.POST("/ping", handler.TestPost)
	v1_r.POST("/family", handler.FamilyPost)

	//r.Run() // listen and serve on 0.0.0.0:8080
	return r
}
