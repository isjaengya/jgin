package schema

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	Db "jgin/api/service"
)

func ValidateUniqFamilyName(fl validator.FieldLevel) bool {
	var i int
	value := fl.Field().String()
	mysqldb := Db.GetMysqlClient()
	err := mysqldb.QueryRow("select count(*) from family where family_name = ?", value).Scan(&i)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	fmt.Println("count --> ", i)
	return i < 1
}
