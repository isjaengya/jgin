package tasks

import (
	"fmt"
	"jgin/api/model"
	ginService "jgin/api/service"
)

func GetUserFirst() (err error) {
	var user model.User
	mysqlClient := ginService.GetMysqlClient()
	err = mysqlClient.Ping()
	fmt.Println(err, "PINGPINGPINGPINGPINGPINGPINGPINGPINGPINGPINGPING")
	err = mysqlClient.QueryRow("select * from user limit 1").Scan(&user.Id, &user.CreateAt, &user.UpdateAt, &user.Uid, &user.FamilyId)
	if err != nil {
		fmt.Println(err.Error(), "没有查询到user")
		return err
	}
	fmt.Println("async async async 查询到user", user)
	return err
}
