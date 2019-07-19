package handler

import (
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/middleware"
)

func GetUserLoginDays(c *gin.Context) {
	/*
		先判断用户今天是否领取奖励，如果没领取的话
		判断一下这个用户今天调用过这个接口没有，没调用过在redis里面set个key来标记，然后在查询db，没有写入，有的话修改
		user_task_status_2019_7_19_uid: 1:在则说明登陆没有领取, 2:则说明领取完毕
	*/

	m := make(map[string]interface{})

	user := middleware.GetUser(c)
	// 判断今天的任务完成没有
	b := user.GetUserTodayLoginF() // 今天登录奖励是否领取
	loginDays := user.GetUserLoginDays()
	if b {
		// 是第一次登录
		m["login_days"] = loginDays
		m["get"] = b

		common.SetOK(c, m)
		return
	}
	// 不是第一个登录, 获取连续登录的天数
	m["login_days"] = loginDays
	m["get"] = b
	common.SetOK(c, m)
	return
}
