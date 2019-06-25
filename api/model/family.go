package model

import (
	"database/sql"
	"fmt"
	Db "tebu_go/api/service"
	"tebu_go/api/util"
)

type Family struct {
	FamilyName string `json:"family_name" binding:"required"`
	MemberName string `json:"member_name" binding:"required"`
	MemberSex string `json:"member_sex" binding:"required"`
	MemberAge string `json:"member_age" binding:"required"`
	MemberCity string `json:"member_city" binding:"required"`
	ChildName string `json:"child_name" binding:"required"`
	ChildSex string `json:"child_sex" binding:"required"`
	ChildAge string `json:"child_age" binding:"required"`
	InvitationCode string `json:"invitation_code"`
}

func (family Family) FindByFamilyCode (code string) (b bool, err error) {
	mysqldb := Db.GetMysqlClient()
	rows, err := mysqldb.Query("select count(*) from family where invitation_code = ?", code)
	if err != nil{
		fmt.Println(err.Error(), "6666666666666666666")
		return true, err
	}
	defer rows.Close()

	var i int32
	if rows.Next() {rows.Scan(&i)}
	return util.CountVerify(i), nil
}

func (family Family) FindByFamilyName () (b bool, err error) {
	mysqldb := Db.GetMysqlClient()
	rows, err := mysqldb.Query("select count(*) from family where family_name = ?", family.FamilyName)
	if err != nil{
		fmt.Println(err.Error(), "333333333333333333")
		return true, err
	}
	defer rows.Close()

	var i int32
	if rows.Next() {rows.Scan(&i)}
	fmt.Println(i, "44444444444444444444")
	//return rows.Next(), nil
	return util.CountVerify(i), nil
}

func (family Family) ChildAdd (tx *sql.Tx) (childId int64) {
	sql := "insert into child (child_name, child_sex, child_age) values (?, ?, ?)"
	// 获取child_id
	child, err := tx.Exec(sql, family.ChildName, family.ChildSex, family.ChildAge)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	childId, err = child.LastInsertId()
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	return childId

}

func (family Family) FamilyAdd () (err error) {
	mysqldb := Db.GetMysqlClient()
	tx, err := mysqldb.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer tx.Rollback()

	childId := family.ChildAdd(tx)

	sql := "insert into family (family_name, member_name, member_sex, member_age, member_city, invitation_code, child_id, run_total, task_stage) values (?,?,?,?,?,?,?,?,?)"
	_family, err := tx.Exec(sql, family.FamilyName, family.MemberName, family.MemberSex, family.MemberAge, family.MemberCity, family.InvitationCode, childId, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	lastId, err := _family.LastInsertId()
	fmt.Println(lastId)
	_ = tx.Commit()
	return
}
