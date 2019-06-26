package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"tebu_go/api/lib/e"
	"time"
)

func SetError(c *gin.Context, code int, err error) {
	obj := gin.H{}
	if err == nil{
		obj["message"] = e.GetMessage(code)
	} else {
		obj["message"] = err.Error()
	}
	obj["errcode"] = code
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