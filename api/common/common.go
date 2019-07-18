package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"jgin/api/lib/e"
	"net/http"
	"time"
)

func SetError(c *gin.Context, msg e.Errmsg, err error) {
	err = GetValidatorError(err)
	obj := gin.H{}
	if err == nil {
		obj["message"] = msg[1]
	} else {
		obj["message"] = err.Error()
	}
	obj["errcode"] = msg[0]
	obj["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, obj)
	return
}

func SetOK(c *gin.Context, obj interface{}) {
	_obj := gin.H{}
	_obj["data"] = obj
	_obj["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, _obj)
	return
}

type ValidateError struct {
	Errors map[string]interface{} `json:"errors"`
}

func GetValidatorError(err error) error {
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	for _, _e := range errs {
		tag := _e.Tag()
		errMsg := e.GetValidateMessage(tag)
		valueStruct := _e.StructField()

		//inputValue := _e.Value() 这里获取的类型不确定，所以没办法直接返回
		//s := fmt.Sprintf("输入:%s, 错误原因:%s", inputValue, errMsg)

		s := fmt.Sprintf("字段名称:%s, 错误原因:%s", valueStruct, errMsg)
		//fmt.Println(_e.ActualTag())  // 这里可以手动放开试一试，对应的是序列化失败中的各种信息
		//fmt.Println(_e.Value())
		//fmt.Println(_e.Type())
		//fmt.Println(_e.Tag())
		//fmt.Println(_e.StructNamespace())
		//fmt.Println(_e.StructField())
		//fmt.Println(_e.Namespace())
		//fmt.Println(_e.Kind())
		//fmt.Println(_e.Param())
		return errors.New(s)
	}
	return err
}
