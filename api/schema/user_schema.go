package schema

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserLoginSchema struct {
	Id  int `json:"id" binding:"required,min=1,max=999999"` // 如果不是整形返回错误信息 您输入的不是数字
	Uid int `json:"uid" binding:"required"`
}

type UserInfoSchema struct {
	Uid string `form:"uid"`
}

func (u *UserLoginSchema) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(u, b)

	if err != nil {
		return err
	}
	return nil
}

type UserQuerySchema struct {
	Uids []string `form:"uid" binding:"required"`
}

func (u *UserQuerySchema) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(u, b)

	if err != nil {
		return err
	}
	return nil
}
