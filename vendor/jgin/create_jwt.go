package main

import (
	"fmt"
	"jgin/api/util"
)

func main() {

	type UserInfo map[string]interface{}

	userInfo := make(UserInfo)

	userInfo["uid"] = "111111"

  tokenString1 := util.CreateJwt(userInfo)
  uid, ok := util.ParseTokenUid(tokenString1)
  if ok {
    fmt.Println(uid, "222222222222")
    }

	fmt.Println(tokenString1, "!111111")
}
