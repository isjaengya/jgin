package handler

import (
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/lib/e"
	"jgin/api/schema"
	"jgin/api/util"
)

func FamilyPost (c *gin.Context) {
	v := schema.FamilySchema{}
	if err := v.Bind(c); err != nil {
		common.SetError(c, e.SELECT_ERROR, err)
		return
	}
	// 5位大写邀请码
	invitationCode := util.RandString(5)
	if b, err := v.FindByFamilyCode(invitationCode); err != nil || b == true{
		common.SetError(c, e.CODE_EXIST, err)
		return
	}

	v.InvitationCode = invitationCode
	if err := v.FamilyAdd(); err != nil{
		common.SetError(c, e.ADD_FAMILY_ERROR, err)
		return
	}

	common.SetOK(c, v)
	return

}
