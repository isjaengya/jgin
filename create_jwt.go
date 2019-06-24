package main

import (
	"fmt"
	"tebu_go/api/lib"
	"tebu_go/api/util"
)

func main(){


  type UserInfo map[string] interface{}

  key := lib.JwtKey
  userInfo := make(UserInfo)

  userInfo["uid"] = "144494"

  tokenString1 := util.CreateJwt(key, userInfo)
  uid, ok := util.ParseTokenUid(tokenString1, key)
  if ok {
    fmt.Println(uid, "222222222222")
    }

  fmt.Println(tokenString1, "!111111")
}
