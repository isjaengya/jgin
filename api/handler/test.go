package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tebu_go/api/middleware"
	"tebu_go/api/schema"
)

func TestPost(c *gin.Context) {

	uid := middleware.GetUid(c)
	fmt.Println(uid, "444444444444444444444")

	var familySchema schema.FamilySchema
	if s := c.ShouldBind(&familySchema); s != nil {
		fmt.Println(familySchema)
		c.JSON(http.StatusOK, gin.H{"msg": s.Error()})
		return
	}
	fmt.Println(familySchema)
	c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
	return
}
