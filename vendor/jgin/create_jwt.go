package main

import (
	"fmt"
	"jgin/api/middleware"
)

func main() {

	type UserInfo map[string]interface{}

	userInfo := make(UserInfo)

	userInfo["uid"] = "111111"

  tokenString1 := middleware.CreateJwt(userInfo)
  uid, ok := middleware.ParseTokenUid(tokenString1)
  if ok {
    fmt.Println(uid, "222222222222")
    }

	fmt.Println(tokenString1, "!111111")
}
