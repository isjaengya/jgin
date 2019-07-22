package model

import (
	"fmt"
	"jgin/api/service"
	"jgin/api/util"
	"time"
)

/* 连续登陆表
CREATE TABLE `user_login_days` (
  `c_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `u_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int(11) NOT NULL,
  `days` int(11) NOT NULL DEFAULT '1',
  `continuous_days` int(11) NOT NULL DEFAULT '1',
  `login_days` int(11) NOT NULL DEFAULT '1',
  `time_str` varchar(128) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `ix_user_login_days_time_str` (`time_str`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type UserLoginDays struct {
	CreateAt       time.Time `json:"create_at" from:"create_at"`
	UpdateAt       time.Time `json:"update_at" from:"update_at"`
	UserId         string    `json:"user_id" from:"user_id"`
	Days           int       `json:"days" form:"days"`
	ContinuousDays int       `json:"continuous_days" from:"continuous_days"`
	LoginDays      int       `json:"login_days" from:"login_days"`
	TimeStr        time.Time `json:"time_str" form:"time_str"`
}

func (u User) FindByUserLoginDays() (userLoginDays UserLoginDays, b bool) {
	//var userLoginDays UserLoginDays
	// 查找用户的的登录记录
	mysqlClient := service.GetMysqlClient()
	if err := mysqlClient.QueryRow("select * from user_login_days where user_id = ?", u.Uid).Scan(&userLoginDays.CreateAt, &userLoginDays.UpdateAt, &userLoginDays.UserId, &userLoginDays.Days, &userLoginDays.ContinuousDays, &userLoginDays.LoginDays, &userLoginDays.TimeStr); err != nil {
		fmt.Println("查询用户登录信息出错, ", err.Error())
		return userLoginDays, false
	}
	return userLoginDays, true
}

func (u User) InsertUserLoginDaysRecord() (b bool) {
	// 插入一条记录
	mysqlClient := service.GetMysqlClient()
	stmt, err := mysqlClient.Prepare("insert into user_login_days (user_id, time_str) value (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	//res, err := stmt.Exec(u.Uid, util.GetTodayStr())
	_, err = stmt.Exec(u.Uid, util.GetTodayStr())
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	//lastId, err := res.LastInsertId()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return false
	//}
	return true
}

func (u *UserLoginDays) Verify() (b bool) {
	// 在这里验证需不需更新连续登录天数
	y1, m1, d1 := u.TimeStr.Date()
	fmt.Println(y1, m1, d1)
	y2, m2, d2 := time.Now().Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		return false
	}
	y2, m2, d2 = time.Now().AddDate(0, 0, -1).Date()
	fmt.Println(y2, m2, d2)
	if y1 == y2 && m1 == m2 && d1 == d2 {
		// 昨天登录过, 连续登陆天数+1
		u.ContinuousDays++
	} else {
		// 连续登录断了，从第一天开始算
		u.ContinuousDays = 1
	}
	u.LoginDays++
	u.TimeStr = time.Now()
	return true
}

func (u UserLoginDays) Update() (b bool) {
	mysqlClient := service.GetMysqlClient()
	stme, err := mysqlClient.Prepare("update user_login_days set days = ?, continuous_days = ?, login_days = ?, time_str = ? where user_id = ?")
	if err != nil {
		fmt.Println(err.Error(), "ssssssssssss")
		return false
	}
	res, err := stme.Exec(u.Days, u.ContinuousDays, u.LoginDays, u.TimeStr, u.UserId)
	if err != nil {
		fmt.Println(err.Error(), "ddddddddddddddd")
		return false
	}
	var rows int64
	rows, err = res.RowsAffected()
	if err != nil {
		fmt.Println(err.Error(), "qqqqqqqqqqq")
		return false
	}
	if rows >= 1 {
		return true
	} else {
		return false
	}

}
