package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"jgin/api/common"
	"jgin/api/lib/e"
	"jgin/api/schema"
)

func TestPost(c *gin.Context) {

	v := schema.FamilySchema{}
	err := v.Bind(c)
	if err != nil{
		common.SetError(c, e.SHOULD_ERROR, err)
		return
	}
	fmt.Println(v)
	c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
	return
}
