package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"min=12,max=15, errmsg:123"`
	Phone  string `validate:"myParam"`
	Age    int    `validate:"min=12,max=15"`
	//CreateAt time.Time `validate:"myParam=this is called param 123123123"`
	CreateAt time.Time `validate:""`
}

func main() {
	address := &Address{
		Street:   "Eavesdown Docks",
		City:     "beijing",
		Planet:   "111",
		Phone:    "132425367",
		Age:      12,
		CreateAt: time.Now(),
	}
	validate := validator.New()
	//自己定义tag标签以及与之对应的处理逻辑
	err := validate.RegisterValidation("myParam", mytimeFunc)
	//if err != nil{
	//    fmt.Println(err.Error())
	//}
	//查看是否符合验证
	err = validate.Struct(address)
	//if err != nil {
	//    s := SetValidatorError(err)
	//    fmt.Println(s)
	//    for k, v := range s.Errors{
	//       fmt.Println(k, v, "dddddddddddddddd")
	//    }
	//}
	fmt.Println(err)
}

func mytimeFunc(fl validator.FieldLevel) bool {
	fmt.Println("FieldName:", fl.FieldName())
	fmt.Println("StructFieldName", fl.StructFieldName())
	fmt.Println("Parm:", fl.Param())
	fmt.Println(fl.Field().String(), "aaaaaaaaaaaaaa")
	return true
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func SetValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		res.Errors[e.Field()] = e.ActualTag()
	}
	return res
}
