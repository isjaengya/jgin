package handler

import (
	"github.com/gin-gonic/gin"
	"tebu_go/api/common"
	"tebu_go/api/lib/e"
	"tebu_go/api/model"
	"tebu_go/api/util"
)

func FamilyPost (c *gin.Context) {
	var family model.Family
	err := c.ShouldBind(&family)
	if err != nil {
		common.SetError(c, e.SHOULD_ERROR, err)
		return
	}

	if b, err := family.FindByFamilyName(); err != nil {
		common.SetError(c, e.FAMILY_EXIST, err)
		return
	} else {
		if b == true{
		common.SetError(c, e.FAMILY_EXIST, err)
		return
		}
	}

	// 5位大写邀请码
	invitationCode := util.RandString(5)
	if b, err := family.FindByFamilyCode(invitationCode); err != nil{
		common.SetError(c, e.CODE_EXIST, err)
		return
	} else {
		if b == true{
		common.SetError(c, e.CODE_EXIST, err)
		return
		}
	}

	family.InvitationCode = invitationCode
	err = family.FamilyAdd()
	if err != nil{
		common.SetError(c, e.ADD_FAMILY_ERROR, err)
		return
	}

	common.SetOK(c, family)
	return

}
