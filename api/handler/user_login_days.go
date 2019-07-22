package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jgin/api/common"
	"jgin/api/lib/e"
	"jgin/api/middleware"
)

func GetUserLoginDays(c *gin.Context) {
	/*
		判断用户今天是否登录接口, 领取连续登录奖励之前的接口
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

func UserReceiveAward(c *gin.Context) {
	/*
		用户领取连续登录奖励接口
		先判断用户今天有没有登录以及是否领取(redis key >= 2)，如果有的话发放给用户奖励，同时更新表中信息
	*/
	user := middleware.GetUser(c)
	b := user.VerifyUserReceiveAward()
	if !b {
		common.SetError(c, e.USER_NOT_LOGIN_TODAY, nil)
		return
	}
	// 发放奖励完成在把user task状态的key incr, 检查mysql里面数据, 没有插入有的话更新
	userLoginDays, ok := user.FindByUserLoginDays()
	if !ok {
		// 添加一条新纪录
		ok = user.InsertUserLoginDaysRecord()
		fmt.Println(ok, "添加用户登录记录是否成功")
		if !ok {
			fmt.Println("添加用户记录失败，请重试。")
			return
		}
	} else {
		// 更新记录
		ok = userLoginDays.Verify()
		if ok {
			if userLoginDays.Update() {
				fmt.Println("更新记录成功")
			}
		}
		fmt.Println(userLoginDays)
	}
	// 插入成功发放奖励, 发放奖励应该是通用接口，接受2个参数，uid以及award_id, 发放任务的接口不应该验证任务是否完成，只给用户增加奖励

}
