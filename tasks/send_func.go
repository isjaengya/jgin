package tasks

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	ginService "jgin/api/service"
)

/* 这里存放所有的异步任务发送函数
   构造signature，在业务逻辑上只调用这里的方法即可
*/
func AsyncHelloWorld(i int32) {
	signature := &tasks.Signature{
		//Name: "hello",
		Name: "hello",
		Args: []tasks.Arg{
			{
				Type:  "int32",
				Value: i,
			},
		},
	}
	asyncResult, err := ginService.GetMachinerty().SendTask(signature)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(asyncResult.GetState().IsSuccess())
	}
}

func AsyncGetUserFirst() {
	signature := &tasks.Signature{
		Name: "get_first_user",
	}
	asyncResult, err := ginService.GetMachinerty().SendTask(signature)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(asyncResult.GetState().IsSuccess())
	}
}
