package model

import (
	"fmt"
	"github.com/francoispqt/gojay"
	"log"
	"tebu_go/api/lib"
	Schema "tebu_go/api/schema"
	"tebu_go/api/service"
	"time"
)

type UnmarshalerJSONObject interface {
	UnmarshalJSONObject(*gojay.Decoder, string) error
	NKeys() int
}

type User struct {
	Id int `json:"id" from:"id"`
	CreateAt int32 `json:"create_at" from:"create_at"`
	UpdateAt int32 `json:"update_at" from:"update_at"`
	Uid string `json:"uid" from:"uid"`
	FamilyId int `json:"family_id" form:"family_id"`
}

func (u *User) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", u.Id)
	enc.StringKey("uid", u.Uid)
	enc.IntKey("family_id", u.FamilyId)
	enc.Int32Key("create_at", u.CreateAt)
}

func (u *User) IsNil() bool {
	return u == nil
}

func (u *User) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
    switch key {
    case "id":
        return dec.Int(&u.Id)
    case "uid":
        return dec.String(&u.Uid)
    case "family_id":
        return dec.Int(&u.FamilyId)
	case "create_at":
		return dec.Int32(&u.CreateAt)

    }
    return nil
}
func (u *User) NKeys() int {
    return 4
}

func GetRedisInfoToUser(uid string) (u *User) {
	user := &User{}
	redisClient := service.GetRedisClient()
	key := lib.UserRedisKey + uid
	s, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("s", s)
	err = gojay.UnmarshalJSONObject([]byte(s), user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func (u *User) SetUserJwtLast10(jwt string) {
	redisClient := service.GetRedisClient()
	key := lib.UserJwtRedisKey + u.Uid
	err := redisClient.Set(key, jwt, 0).Err()
	if err != nil {
		fmt.Println("set user jwt last 10 error, ", err.Error())
	}
}

func (u *User) UpdateRedisCache() {
	json := u.ToJson()
	redisClient := service.GetRedisClient()

	key := lib.UserRedisKey + u.Uid
	err := redisClient.Set(key, json, time.Second*3600).Err()
	if err != nil {
		fmt.Println("set user info error, ", err.Error())
	}
}

func (u *User) ToJson() (s string) {
	b, err := gojay.MarshalJSONObject(u)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func VerifyUserLogin(schema Schema.UserLoginSchema) (user User, b bool) {
	id := schema.Id
	uid := schema.Uid

	mysqldb := service.GetMysqlClient()
	err := mysqldb.QueryRow("select * from user where id = ? and uid = ?", id, uid).Scan(&user.Id, &user.CreateAt, &user.UpdateAt, &user.Uid, &user.FamilyId)
	if err != nil{
		fmt.Println(err.Error(), "没有查询到user")
		return user, false
	} else {
		fmt.Println("查询到user", user)
		return user, true
	}
}