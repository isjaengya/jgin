package schema

import (
	"database/sql"
	"fmt"
	Db "tebu_go/api/service"
)

type FamilySchema struct {
	FamilyName string `json:"family_name" binding:"required"`
	MemberName string `json:"member_name" binding:"required"`
	MemberSex string `json:"member_sex" binding:"required"`
	MemberAge string `json:"member_age" binding:"required"`
	MemberCity string `json:"member_city" binding:"required"`
	ChildName string `json:"child_name" binding:"required"`
	ChildSex string `json:"child_sex" binding:"required"`
	ChildAge string `json:"child_age" binding:"required"`
}

func (family FamilySchema) ChildAdd (tx *sql.Tx) (childId int64) {
	sql := fmt.Sprintf("insert into child (child_name, child_sex, child_age) values (\"%s\", \"%s\", \"%s\")", family.ChildName, family.ChildSex, family.ChildAge)
	fmt.Println(sql)

		// 获取child_id
	child, err := tx.Exec(sql)
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

func (family FamilySchema) FamilyAdd () (err error) {
	tx, err := Db.Mysqldb.Begin()
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	defer tx.Rollback()

	childId := family.ChildAdd(tx)
	fmt.Println(childId)

	sql := "insert into family (family_name, member_name, member_sex, member_age, member_city, invitation_code, child_id, run_total, task_stage, uid) values ()"
	fmt.Println(sql)
	err = tx.Commit()
	return
}