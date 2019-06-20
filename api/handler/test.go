package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tebu_go/api/schema"
)

func TestPost(c *gin.Context) {

	//age := c.Query("member_age")
	//fmt.Println(age, "cccccccccccccccc")

	var familySchema schema.FamilySchema
	s := c.ShouldBind(&familySchema)
	if s != nil{
		fmt.Println(familySchema)
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": s.Error()})
		return
	}
	fmt.Println(familySchema)
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
