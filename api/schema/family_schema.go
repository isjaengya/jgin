package schema

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	Db "jgin/api/service"
	"jgin/api/util"
)

type FamilySchema struct {
	FamilyName string `json:"family_name" binding:"required,family_name-uniq"`
	MemberName string `json:"member_name" binding:"required"`
	MemberSex string `json:"member_sex" binding:"required"`
	MemberAge string `json:"member_age" binding:"required"`
	MemberCity string `json:"member_city" binding:"required"`
	ChildName string `json:"child_name" binding:"required"`
	ChildSex string `json:"child_sex" binding:"required"`
	ChildAge string `json:"child_age" binding:"required"`
	InvitationCode string `json:"invitation_code"`
}

func (f *FamilySchema) Bind (c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(f, b)

	if err != nil {
		return err
	}
	return nil
}

func (f FamilySchema) FindByFamilyCode (code string) (b bool, err error) {
	var i int32
	mysqlDb := Db.GetMysqlClient()
	err = mysqlDb.QueryRow("select count(*) from family where invitation_code = ?", code).Scan(&i)
	if err != nil {
		return true, err
	}
	return util.CountVerify(i), nil
}

func (f FamilySchema) FindByFamilyName () (b bool, err error) {
	mysqlDb := Db.GetMysqlClient()
	var i int32
	err = mysqlDb.QueryRow("select count(*) from family where family_name = ?", f.FamilyName).Scan(&i)
	if err != nil{
		return true, err
	}
	return util.CountVerify(i), nil
}

func (f FamilySchema) ChildAdd (tx *sql.Tx) (childId int64) {
	_sql := "insert into child (child_name, child_sex, child_age) values (?, ?, ?)"
	// 获取child_id
	child, err := tx.Exec(_sql, f.ChildName, f.ChildSex, f.ChildAge)
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

func (f FamilySchema) FamilyAdd () (err error) {
	mysqlDb := Db.GetMysqlClient()
	tx, err := mysqlDb.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer tx.Rollback()

	childId := f.ChildAdd(tx)

	_sql := "insert into family (family_name, member_name, member_sex, member_age, member_city, invitation_code, child_id, run_total, task_stage) values (?,?,?,?,?,?,?,?,?)"
	_family, err := tx.Exec(_sql, f.FamilyName, f.MemberName, f.MemberSex, f.MemberAge, f.MemberCity, f.InvitationCode, childId, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	lastId, err := _family.LastInsertId()
	fmt.Println(lastId)
	_ = tx.Commit()
	return
}

