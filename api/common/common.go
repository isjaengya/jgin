package common

import (
	"github.com/gin-gonic/gin"
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